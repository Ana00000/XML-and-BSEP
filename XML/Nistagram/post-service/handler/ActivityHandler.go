package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	gomail "gopkg.in/mail.v2"
	"net/http"
	"os"
	"time"
)

type ActivityHandler struct {
	Service * service.ActivityService
	SinglePostService * service.SinglePostService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response := Request(reqUrlAuth, ExtractToken(r))
	if response.StatusCode == 401 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "CRACT467",
			"timestamp": time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-activity-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization, ExtractToken(r))
	if res.StatusCode == 403 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "CRACT467",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}
	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "CRACT467",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var activityDTO dto.ActivityDTO
	err := json.NewDecoder(r.Body).Decode(&activityDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "CRACT467",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to ActivityDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var activityFromDatabase = handler.Service.FindByPostIDAndUserID(activityDTO.PostID,activityDTO.UserID)
	if activityFromDatabase!=nil{
		activityDTO.ID=activityFromDatabase.ID
		err := handler.Service.UpdateActivity(&activityDTO)
		if err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status":    "failure",
				"location":  "ActivityHandler",
				"action":    "CRACT467",
				"timestamp": time.Now().String(),
			}).Error("Failed update activity!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		activityIDJson, _ := json.Marshal(activityDTO.ID)
		w.Write(activityIDJson)
		w.WriteHeader(http.StatusOK)

		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "CRACT467",
			"timestamp": time.Now().String(),
		}).Info("Successfully updated activity!")

		w.Header().Set("Content-Type", "application/json")
		return
	}

	activity := model.Activity{
		ID:          uuid.UUID{},
		PostID:      activityDTO.PostID,
		UserID:      activityDTO.UserID,
		LikedStatus: activityDTO.LikedStatus,
		IsFavorite:  activityDTO.IsFavorite,
	}

	err = handler.Service.CreateActivity(&activity)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "CRACT467",
			"timestamp": time.Now().String(),
		}).Error("Failed creating activity!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	activityIDJson, _ := json.Marshal(activity.ID)
	w.Write(activityIDJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ActivityHandler",
		"action":    "CRACT467",
		"timestamp": time.Now().String(),
	}).Info("Successfully created activity!")


	var userId = handler.SinglePostService.FindOwnerOfPost(activityDTO.PostID)
	var user dto.ClassicUserDTO
	reqUrlUser := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), userId)
	err = getJson(reqUrlUser, &user)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "UPACT472",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find user by id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var profileSettings dto.ProfileSettingsDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), user.ID)
	err = getJson(reqUrl, &profileSettings)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "UPACT472",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find profile settings by user id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if profileSettings.LikesNotifications == "ALL_NOTIFICATIONS"{
		//SEND EMAIL NOTIFICATION
		handler.SendNotificationMail(user.Email, activityDTO)
	}else if profileSettings.LikesNotifications == "FRIENDS_NOTIFICATION"{
		//check if senderUser is friend
		var followings []dto.ClassicUserFollowingsFullDTO
		reqUrl := fmt.Sprintf("http://%s:%s/find_all_valid_followings_for_user/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), user.ID)
		err = getJson(reqUrl, &followings)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "ActivityHandler",
				"action":   "UPACT472",
				"timestamp":   time.Now().String(),
			}).Error("Failed to find followings for user!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

		for  i:=0; i < len(followings); i++{
			if followings[i].FollowingUserId == activityDTO.UserID {
				//SEND EMAIL NOTIFICATION
				handler.SendNotificationMail(user.Email, activityDTO)
			}
		}
	}


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ActivityHandler) FindAllLikesForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllLikesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "FALFP468",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all likes for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
		return
	}

	handler.LogError.WithFields(logrus.Fields{
		"status":    "failure",
		"location":  "ActivityHandler",
		"action":    "FALFP468",
		"timestamp": time.Now().String(),
	}).Error("Likes for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllDislikesForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllDislikesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "FADFP469",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all dislikes for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status":    "failure",
		"location":  "ActivityHandler",
		"action":    "FALFP469",
		"timestamp": time.Now().String(),
	}).Error("Dislikes for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllFavoritesForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllFavoritesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "FAFFP470",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all favorites for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status":    "failure",
		"location":  "ActivityHandler",
		"action":    "FAFFP470",
		"timestamp": time.Now().String(),
	}).Error("Favorites for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllActivitiesForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response := Request(reqUrlAuth, ExtractToken(r))
	if response.StatusCode == 401 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FAAFP471",
			"timestamp": time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAuthorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-activities-for-post-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAuthorization, ExtractToken(r))
	if res.StatusCode == 403 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FAAFP471",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "FAAFP471",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	id := r.URL.Query().Get("id")

	activities := handler.Service.FindAllActivitiesForPost(uuid.MustParse(id))
	activitiesJson, _ := json.Marshal(activities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "FAAFP471",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all activities for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status":    "failure",
		"location":  "ActivityHandler",
		"action":    "FAAFP471",
		"timestamp": time.Now().String(),
	}).Error("Activites for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) FindAllLikedPostsByUserId(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response := Request(reqUrlAuth, ExtractToken(r))
	if response.StatusCode == 401 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ActivityHandler",
			"action": "FindAllLikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-liked-post-by-user-id-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization, ExtractToken(r))
	if res.StatusCode == 403 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ActivityHandler",
			"action": "FindAllLikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	userId := r.URL.Query().Get("id")

	allLikedPostActivities := handler.Service.FindAllLikedPostsByUserId(uuid.MustParse(userId))

	allLikedPosts := handler.SinglePostService.FindAllPostsByIds(returnListIDs(allLikedPostActivities))

	var posts = convertListSinglePostsToSinglePostsDTO(allLikedPosts)
	//finds all conents
	//var contents = handler.PostContentService.FindAllContentsForPosts(posts)
	reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_posts/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(posts)
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllLikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Error("Failed to find all contents for posts!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SinglePostContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllLikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to SinglePostContentDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//finds all locations
	//var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_posts/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(posts)
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllLikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Error("Failed to find all locations for posts!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllLikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts) //treba izmjeniti

	//find all tags
	//var tags = handler.PostTagPostsService.FindAllTagsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post_tag_posts/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(posts)
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllLikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Error("Failed to find all tags for posts!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostTagPostsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllLikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to PostTagPostsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//creates a list of dtos
	var postsDTOS = handler.CreatePostsDTOList(convertSinglePostsDTOToListSinglePosts(posts), contents, locations, tags)

	activitiesJson, _ := json.Marshal(postsDTOS)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "FindAllLikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all posts this user liked!")
		w.Write(activitiesJson)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	handler.LogError.WithFields(logrus.Fields{
		"status":    "failure",
		"location":  "ActivityHandler",
		"action":    "FindAllLikedPostByUserId",
		"timestamp": time.Now().String(),
	}).Error("No liked post were found for this user!")
	w.WriteHeader(http.StatusBadRequest)
}

func returnListIDs(list []model.Activity) []uuid.UUID{
	var retVal []uuid.UUID
	for i := 0; i < len(list); i++ {
		retVal = append(retVal, list[i].PostID)
	}
	return retVal
}

func (handler *ActivityHandler) FindAllDislikedPostsByUserId(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response := Request(reqUrlAuth, ExtractToken(r))
	if response.StatusCode == 401 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ActivityHandler",
			"action": "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}
	reqUrlAuthorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-disliked-post-by-user-id-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAuthorization, ExtractToken(r))
	if res.StatusCode == 403 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ActivityHandler",
			"action": "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 403
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	userId := r.URL.Query().Get("id")

	allDislikedPostActivities := handler.Service.FindAllDislikedPostsByUserId(uuid.MustParse(userId))

	allDislikedPosts := handler.SinglePostService.FindAllPostsByIds(returnListIDs(allDislikedPostActivities))

	var posts = convertListSinglePostsToSinglePostsDTO(allDislikedPosts)

	reqUrl := fmt.Sprintf("http://%s:%s/find_all_contents_for_posts/", os.Getenv("CONTENT_SERVICE_DOMAIN"), os.Getenv("CONTENT_SERVICE_PORT"))
	jsonValidPostsDTO, _ := json.Marshal(posts)
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonValidPostsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("Failed to find all contents for posts!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	//defer resp.Body.Close() mozda treba dodati
	var contents []dto.SinglePostContentDTO
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to SinglePostContentDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//finds all locations
	//var locations = handler.LocationService.FindAllLocationsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_locations_for_posts/", os.Getenv("LOCATION_SERVICE_DOMAIN"), os.Getenv("LOCATION_SERVICE_PORT"))
	jsonLocationsDTO, _ := json.Marshal(posts)
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonLocationsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("Failed to find all locations for posts!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var locations []dto.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to LocationDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//var tags = handler.PostTagPostsService.FindAllTagsForPostsTagPosts(posts) //treba izmjeniti

	//find all tags
	//var tags = handler.PostTagPostsService.FindAllTagsForPosts(posts)
	reqUrl = fmt.Sprintf("http://%s:%s/find_all_tags_for_post_tag_posts/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonTagsDTO, _ := json.Marshal(posts)
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonTagsDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("Failed to find all tags for posts!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var tags []dto.PostTagPostsDTO
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to PostTagPostsDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}

	//creates a list of dtos
	var postsDTOS = handler.CreatePostsDTOList(convertSinglePostsDTOToListSinglePosts(posts), contents, locations, tags)

	activitiesJson, _ := json.Marshal(postsDTOS)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "FindAllDislikedPostsByUserId",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all posts this user disliked!")
		w.Write(activitiesJson)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	handler.LogError.WithFields(logrus.Fields{
		"status":    "failure",
		"location":  "ActivityHandler",
		"action":    "FindAllDislikedPostByUserId",
		"timestamp": time.Now().String(),
	}).Error("No disliked post were found for this user!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response := Request(reqUrlAuth, ExtractToken(r))
	if response.StatusCode == 401 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "UPACT472",
			"timestamp": time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-update-activity-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization, ExtractToken(r))
	if res.StatusCode == 403 {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "UPACT472",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var activityDTO dto.ActivityDTO

	err := json.NewDecoder(r.Body).Decode(&activityDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "UPACT472",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to ActivityDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.UpdateActivity(&activityDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ActivityHandler",
			"action":    "UPACT472",
			"timestamp": time.Now().String(),
		}).Error("Activity not updated!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ActivityHandler",
		"action":    "UPACT472",
		"timestamp": time.Now().String(),
	}).Info("Successfully updated activity!")

	//GET USERID FROM POSTID FOR ACTIVITYNOTIFICATION
	/*var userId uuid.UUID
	reqUrl := fmt.Sprintf("http://%s:%s/find_owner_of_post/%s", os.Getenv("POST_SERVICE_DOMAIN"), os.Getenv("POST_SERVICE_DOMAIN"), activityDTO.PostID)
	err = getJson(reqUrl, &userId)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "UPACT472",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find owner of the post!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}*/
	var userId = handler.SinglePostService.FindOwnerOfPost(activityDTO.PostID)
	var user dto.ClassicUserDTO
	reqUrlUser := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), userId)
	err = getJson(reqUrlUser, &user)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "UPACT472",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find user by id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var profileSettings dto.ProfileSettingsDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), user.ID)
	err = getJson(reqUrl, &profileSettings)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "UPACT472",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find profile settings by user id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if profileSettings.LikesNotifications == "ALL_NOTIFICATIONS"{
		//SEND EMAIL NOTIFICATION
		handler.SendNotificationMail(user.Email, activityDTO)
	}else if profileSettings.LikesNotifications == "FRIENDS_NOTIFICATION"{
		//check if senderUser is friend
		var followings []dto.ClassicUserFollowingsFullDTO
		reqUrl := fmt.Sprintf("http://%s:%s/find_all_valid_followings_for_user/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), user.ID)
		err = getJson(reqUrl, &followings)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "ActivityHandler",
				"action":   "UPACT472",
				"timestamp":   time.Now().String(),
			}).Error("Failed to find followings for user!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

		for  i:=0; i < len(followings); i++{
			if followings[i].FollowingUserId == activityDTO.UserID {
				//SEND EMAIL NOTIFICATION
				handler.SendNotificationMail(user.Email, activityDTO)
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ActivityHandler) SendNotificationMail(email string, activity dto.ActivityDTO) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "xml.ftn.uns@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Confirmation mail")

	var user dto.ClassicUserDTO
	reqUrlUser := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), activity.UserID)
	err := getJson(reqUrlUser, &user)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "SEDCONFMAIL777",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find user by id!")
		panic(err)
	}

	// Set E-Mail body. You can set plain text or html with text/html
	text := user.FirstName + " " + user.LastName + "\n\n\n reacted to your post!\n\nhttps://localhost:8081/postById/" + activity.PostID.String()
	if activity.LikedStatus == 0 {
		text += "\n\n\nPost is currently liked"
	} else if activity.LikedStatus == 1 {
		text += "\n\n\nPost is currently disliked"
	}

	if activity.IsFavorite == true {
		text += " and favorited."
	} else {
		text += "."
	}

	text += "\n\n\nBest regards,\nTim25"
	m.SetBody("text/plain", text)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "xml.ftn.uns@gmail.com", "XMLFTNUNS1")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		//fmt.Println(err)
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "SEDCONFMAIL777",
			"timestamp":   time.Now().String(),
		}).Error("Failed sending email with confirmation token!")
		panic(err)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ActivityHandler",
		"action":   "SEDCONFMAIL777",
		"timestamp":   time.Now().String(),
	}).Info("Successfully sent email with confirmation token!")
}

func (handler *ActivityHandler) CreatePostsDTOList(posts []model.SinglePost, contents []dto.SinglePostContentDTO, locations []dto.LocationDTO, tags []dto.PostTagPostsDTO) []dto.SelectedPostDTO {
	var listOfPostsDTOs []dto.SelectedPostDTO

	for i := 0; i < len(posts); i++ {
		var postDTO dto.SelectedPostDTO
		postDTO.PostId = posts[i].ID
		postDTO.Description = posts[i].Description
		postDTO.CreationDate = posts[i].CreationDate
		postDTO.UserId = posts[i].UserID

		for j := 0; j < len(contents); j++ {
			if contents[j].SinglePostId == posts[i].ID {
				postDTO.Path = contents[j].Path
				postDTO.Type = contents[j].Type
			}
		}

		for k := 0; k < len(locations); k++ {
			if locations[k].ID == posts[i].LocationId {
				postDTO.LocationId = locations[k].ID
				postDTO.City = locations[k].City
				postDTO.Country = locations[k].Country
				postDTO.StreetName = locations[k].StreetName
				postDTO.StreetNumber = locations[k].StreetNumber
			}
		}

		var listOfTags []string
		for p := 0; p < len(tags); p++ {
			if tags[p].PostId == posts[i].ID {
				var returnValueTagName ReturnValueString
				reqUrl := fmt.Sprintf("http://%s:%s/get_tag_name_by_id/%s", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"), tags[p].TagId)
				err := getJson(reqUrl, &returnValueTagName)
				if err != nil {
					handler.LogError.WithFields(logrus.Fields{
						"status":    "failure",
						"location":  "ActivityHandler",
						"action":    "CLDTO621",
						"timestamp": time.Now().String(),
					}).Error("Failed to get tag name by id!")
					panic(err)
				}
				listOfTags = append(listOfTags, returnValueTagName.ReturnValue)
			}
		}

		postDTO.Tags = listOfTags

		listOfPostsDTOs = append(listOfPostsDTOs, postDTO)

	}

	return listOfPostsDTOs

}