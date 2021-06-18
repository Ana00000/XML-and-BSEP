package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
	"os"
	"time"
)

type ActivityHandler struct {
	Service  *service.ActivityService
	LogInfo  *logrus.Logger
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
	userId := r.URL.Query().Get("user_id")

	allLikedPostActivities := handler.Service.FindAllLikedPostsByUserId(uuid.MustParse(userId))
	activitiesJson, _ := json.Marshal(allLikedPostActivities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "FindAllLikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all posts this user liked!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
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

	userId := r.URL.Query().Get("user_id")

	allDislikedPostActivities := handler.Service.FindAllDislikedPostsByUserId(uuid.MustParse(userId))
	activitiesJson, _ := json.Marshal(allDislikedPostActivities)
	if activitiesJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ActivityHandler",
			"action":    "FindAllDislikedPostByUserId",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all posts this user disliked!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(activitiesJson)
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

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ActivityHandler",
			"action":   "UPACT472",
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
