package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type ProfileSettingsHandler struct {
	Service   *service.ProfileSettingsService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *ProfileSettingsHandler) CreateProfileSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]

	profileSettings := model.ProfileSettings{
		ID:                  uuid.UUID{},
		UserId:              uuid.MustParse(userId),
		UserVisibility:      model.PUBLIC_VISIBILITY,
		MessageApprovalType: model.PUBLIC,
		IsPostTaggable:      true,
		IsStoryTaggable:     true,
		IsCommentTaggable:   true,
		LikesNotifications: model.ALL_NOTIFICATIONS,
		CommentsNotifications: model.ALL_NOTIFICATIONS,
		MessagesNotifications: model.ALL_NOTIFICATIONS,
	}

	if err := handler.Service.CreateProfileSettings(&profileSettings); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsHandler",
			"action":    "CRPROFSETTINGS2712",
			"timestamp": time.Now().String(),
		}).Error("Failed creating profile settings!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsHandler",
		"action":    "CRPROFSETTINGS2712",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ProfileSettingsHandler) FindProfileSettingWithIDByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]

	var profileSettings = handler.Service.FindProfileSettingByUserId(uuid.MustParse(userId))
	if profileSettings == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsHandler",
			"action":    "FindProfileSettingWithIDByUserId",
			"timestamp": time.Now().String(),
		}).Error("Profile setting not found!")
		fmt.Println("Profile setting not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userVisibility := ""
	if profileSettings.UserVisibility == model.PRIVATE_VISIBILITY {
		userVisibility = "PRIVATE_VISIBILITY"
	} else if profileSettings.UserVisibility == model.PUBLIC_VISIBILITY {
		userVisibility = "PUBLIC_VISIBILITY"
	}

	messageApprovalType := ""
	if profileSettings.MessageApprovalType == model.PUBLIC {
		messageApprovalType = "PUBLIC"
	} else if profileSettings.MessageApprovalType == model.FRIENDS_ONLY {
		messageApprovalType = "FRIENDS_ONLY"
	}

	likesNotifications :=""
	if  profileSettings.LikesNotifications == model.ALL_NOTIFICATIONS{
		likesNotifications = "ALL_NOTIFICATIONS"
	}else if  profileSettings.LikesNotifications == model.FRIENDS_NOTIFICATIONS{
		likesNotifications = "FRIENDS_NOTIFICATIONS"
	}else if profileSettings.LikesNotifications == model.NONE {
		likesNotifications = "NONE"
	}

	commentsNotifications :=""
	if  profileSettings.CommentsNotifications == model.ALL_NOTIFICATIONS{
		commentsNotifications = "ALL_NOTIFICATIONS"
	}else if  profileSettings.CommentsNotifications == model.FRIENDS_NOTIFICATIONS{
		commentsNotifications = "FRIENDS_NOTIFICATIONS"
	}else if profileSettings.CommentsNotifications == model.NONE {
		commentsNotifications = "NONE"
	}

	messagesNotifications :=""
	if  profileSettings.MessagesNotifications == model.ALL_NOTIFICATIONS{
		messagesNotifications = "ALL_NOTIFICATIONS"
	}else if  profileSettings.MessagesNotifications == model.FRIENDS_NOTIFICATIONS{
		messagesNotifications = "FRIENDS_NOTIFICATIONS"
	}else if profileSettings.MessagesNotifications == model.NONE {
		messagesNotifications = "NONE"
	}



	var profileSettingsDTO = dto.ProfileSettingsFullDTO{
		Id:                    profileSettings.ID,
		UserId:                profileSettings.UserId,
		UserVisibility:        userVisibility,
		MessageApprovalType:   messageApprovalType,
		IsPostTaggable:        profileSettings.IsPostTaggable,
		IsStoryTaggable:       profileSettings.IsStoryTaggable,
		IsCommentTaggable:     profileSettings.IsCommentTaggable,
		LikesNotifications:    likesNotifications,
		CommentsNotifications: commentsNotifications,
		MessagesNotifications: messagesNotifications,
	}

	profileSettingsDTOJson, _ := json.Marshal(profileSettingsDTO)
	if profileSettingsDTOJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ProfileSettingsHandler",
			"action":    "FindProfileSettingWithIDByUserId",
			"timestamp": time.Now().String(),
		}).Info("Successfully found profile setting by user id!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(profileSettingsDTOJson)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func (handler *ProfileSettingsHandler) FindProfileSettingByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]

	var profileSettings = handler.Service.FindProfileSettingByUserId(uuid.MustParse(userId))
	if profileSettings == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsHandler",
			"action":    "FINDPROFSETTINGSBYUSID1411",
			"timestamp": time.Now().String(),
		}).Error("Profile setting not found!")
		fmt.Println("Profile setting not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userVisibility := ""
	if profileSettings.UserVisibility == model.PRIVATE_VISIBILITY {
		userVisibility = "PRIVATE_VISIBILITY"
	} else if profileSettings.UserVisibility == model.PUBLIC_VISIBILITY {
		userVisibility = "PUBLIC_VISIBILITY"
	}

	messageApprovalType := ""
	if profileSettings.MessageApprovalType == model.PUBLIC {
		messageApprovalType = "PUBLIC"
	} else if profileSettings.MessageApprovalType == model.FRIENDS_ONLY {
		messageApprovalType = "FRIENDS_ONLY"
	}

	likesNotifications :=""
	if  profileSettings.LikesNotifications == model.ALL_NOTIFICATIONS{
		likesNotifications = "ALL_NOTIFICATIONS"
	}else if  profileSettings.LikesNotifications == model.FRIENDS_NOTIFICATIONS{
			likesNotifications = "FRIENDS_NOTIFICATIONS"
	}else if profileSettings.LikesNotifications == model.NONE {
		likesNotifications = "NONE"
	}

	commentsNotifications :=""
	if  profileSettings.CommentsNotifications == model.ALL_NOTIFICATIONS{
		commentsNotifications = "ALL_NOTIFICATIONS"
	}else if  profileSettings.CommentsNotifications == model.FRIENDS_NOTIFICATIONS{
		commentsNotifications = "FRIENDS_NOTIFICATIONS"
	}else if profileSettings.CommentsNotifications == model.NONE {
		commentsNotifications = "NONE"
	}

	messagesNotifications :=""
	if  profileSettings.MessagesNotifications == model.ALL_NOTIFICATIONS{
		messagesNotifications = "ALL_NOTIFICATIONS"
	}else if  profileSettings.MessagesNotifications == model.FRIENDS_NOTIFICATIONS{
		messagesNotifications = "FRIENDS_NOTIFICATIONS"
	}else if profileSettings.MessagesNotifications == model.NONE {
		messagesNotifications = "NONE"
	}



	var profileSettingsDTO = dto.ProfileSettingsDTO{
		UserId:              profileSettings.UserId,
		UserVisibility:      userVisibility,
		MessageApprovalType: messageApprovalType,
		IsPostTaggable:      profileSettings.IsPostTaggable,
		IsStoryTaggable:     profileSettings.IsStoryTaggable,
		IsCommentTaggable:   profileSettings.IsCommentTaggable,
		LikesNotifications: likesNotifications,
		CommentsNotifications: commentsNotifications,
		MessagesNotifications: messagesNotifications,

	}

	profileSettingsDTOJson, _ := json.Marshal(profileSettingsDTO)
	if profileSettingsDTOJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ProfileSettingsHandler",
			"action":    "FINDPROFSETTINGSBYUSID1411",
			"timestamp": time.Now().String(),
		}).Info("Successfully found profile setting by user id!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(profileSettingsDTOJson)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func (handler *ProfileSettingsHandler) FindProfileSettingsForPublicUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	var profileSettings = handler.Service.FindAllProfileSettingsForPublicUsers()
	if profileSettings == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsHandler",
			"action":    "FIDPROFSETTINGSFORPUBUS0906",
			"timestamp": time.Now().String(),
		}).Error("Profile settings for public users not found!")
		fmt.Println("Profile settings for public users not found!")
		w.WriteHeader(http.StatusNotFound)
	}

	dataJson, _ := json.Marshal(convertListUUIDToListData(profileSettings))
	if dataJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ProfileSettingsHandler",
			"action":    "FIDPROFSETTINGSFORPUBUS0906",
			"timestamp": time.Now().String(),
		}).Info("Successfully found profile settings for public users!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
	}
	w.WriteHeader(http.StatusNotFound)
}

func (handler *ProfileSettingsHandler) FindAllPublicUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var classicUsersDTO []dto.ClassicUserDTO
	if err := json.NewDecoder(r.Body).Decode(&classicUsersDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsHandler",
			"action":    "FIDALLPUBUS3110",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to ClassicUsersDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}


	var classicUsers = handler.Service.FindAllPublicUsers(classicUsersDTO)
	dataJson, _ := json.Marshal(classicUsers)
	if dataJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ProfileSettingsHandler",
			"action":    "FIDALLPUBUS3110",
			"timestamp": time.Now().String(),
		}).Info("Successfully found profile all public users!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type Data struct {
	Uuid uuid.UUID
}

func convertListUUIDToListData(uuids []uuid.UUID) []Data {
	var datas []Data
	for i := 0; i < len(uuids); i++ {
		datas = append(datas, Data{Uuid: uuids[i]})
	}
	return datas
}

func (handler *ProfileSettingsHandler) UpdateProfileSettings(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var profileSettingsDTO dto.ProfileSettingsDTO

	if err := json.NewDecoder(r.Body).Decode(&profileSettingsDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ProfileSettingsHandler",
			"action":   "UPPRSEO777",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ProfileSettingsDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	err := handler.Service.UpdateProfileSettings(&profileSettingsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ProfileSettingsHandler",
			"action":   "UPPRSEO777",
			"timestamp":   time.Now().String(),
		}).Error("Failed updating profile settings!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ProfileSettingsHandler",
		"action":   "UPPRSEO777",
		"timestamp":   time.Now().String(),
	}).Info("Successfully updated profile settings!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


//FIND ALL USERS WITH POST NOTIFICATIONS SET FOR USER
func (handler *ProfileSettingsHandler) FindAllUsersForPostNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]

	var classicUsersIds = handler.Service.FindAllUsersForPostNotifications(uuid.MustParse(userId))
	var classicUsers []dto.ClassicUserDTO
	var user dto.ClassicUserDTO
	for i := 0; i < len(classicUsersIds); i++ {
		reqUrl := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), classicUsersIds[i])
		err := getJson(reqUrl, &user)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "ProfileSettingsHandler",
				"action":   "FIALUSFOPONO455",
				"timestamp":   time.Now().String(),
			}).Error("Failed to find user!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		classicUsers = append(classicUsers, user)
	}
	dataJson, _ := json.Marshal(classicUsers)
	if dataJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ProfileSettingsHandler",
			"action":    "FindAllUsersForPostNotifications",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all users for post notifications!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

//FIND ALL USERS WITH POST ALBUM NOTIFICATIONS SET FOR USER
func (handler *ProfileSettingsHandler) FindAllUsersForPostAlbumNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]

	var classicUsersIds = handler.Service.FindAllUsersForPostAlbumNotifications(uuid.MustParse(userId))
	var classicUsers []dto.ClassicUserDTO
	var user dto.ClassicUserDTO
	for i := 0; i < len(classicUsersIds); i++ {
		reqUrl := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), classicUsersIds[i])
		err := getJson(reqUrl, &user)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "ProfileSettingsHandler",
				"action":   "FIALUSFOPOALNO674",
				"timestamp":   time.Now().String(),
			}).Error("Failed to find user!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		classicUsers = append(classicUsers, user)
	}
	dataJson, _ := json.Marshal(classicUsers)
	if dataJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ProfileSettingsHandler",
			"action":    "FIALUSFOPOALNO674",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all users for post album notifications!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

//FIND ALL USERS WITH STORY NOTIFICATIONS SET FOR USER
func (handler *ProfileSettingsHandler) FindAllUsersForStoryNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]

	var classicUsersIds = handler.Service.FindAllUsersForStoryNotifications(uuid.MustParse(userId))
	var classicUsers []dto.ClassicUserDTO
	var user dto.ClassicUserDTO
	for i := 0; i < len(classicUsersIds); i++ {
		reqUrl := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), classicUsersIds[i])
		err := getJson(reqUrl, &user)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "ProfileSettingsHandler",
				"action":   "FIALUSFOSTNO787",
				"timestamp":   time.Now().String(),
			}).Error("Failed to find user!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		classicUsers = append(classicUsers, user)
	}
	dataJson, _ := json.Marshal(classicUsers)
	if dataJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ProfileSettingsHandler",
			"action":    "FIALUSFOSTNO787",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all users for story notifications!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

//FIND ALL USERS WITH STORY ALBUM NOTIFICATIONS SET FOR USER
func (handler *ProfileSettingsHandler) FindAllUsersForStoryAlbumNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]

	var classicUsersIds = handler.Service.FindAllUsersForStoryAlbumNotifications(uuid.MustParse(userId))
	var classicUsers []dto.ClassicUserDTO
	var user dto.ClassicUserDTO
	for i := 0; i < len(classicUsersIds); i++ {
		reqUrl := fmt.Sprintf("http://%s:%s/get_user_by_id?id=%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), classicUsersIds[i])
		err := getJson(reqUrl, &user)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "ProfileSettingsHandler",
				"action":   "FIALUSFOSTALNO793",
				"timestamp":   time.Now().String(),
			}).Error("Failed to find user!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		classicUsers = append(classicUsers, user)
	}
	dataJson, _ := json.Marshal(classicUsers)
	if dataJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status":    "success",
			"location":  "ProfileSettingsHandler",
			"action":    "FIALUSFOSTALNO793",
			"timestamp": time.Now().String(),
		}).Info("Successfully found all users for story album notifications!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}


func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}