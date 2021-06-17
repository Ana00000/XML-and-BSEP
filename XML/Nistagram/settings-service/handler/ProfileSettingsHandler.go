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
	_ "strconv"
	"time"
)

type ProfileSettingsHandler struct {
	Service   *service.ProfileSettingsService
	ProfileSettingsMutedProfilesHandler * service.ProfileSettingsMutedProfilesService
	ProfileSettingsBlockedProfilesHandler * service.ProfileSettingsBlockedProfilesService
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

	var profileSettingsDTO = dto.ProfileSettingsDTO{
		UserId:              profileSettings.UserId,
		UserVisibility:      userVisibility,
		MessageApprovalType: messageApprovalType,
		IsPostTaggable:      profileSettings.IsPostTaggable,
		IsStoryTaggable:     profileSettings.IsStoryTaggable,
		IsCommentTaggable:   profileSettings.IsCommentTaggable,
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
	/*if err := handler.Validator.Struct(&classicUsersDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsHandler",
			"action":    "FIDALLPUBUS3110",
			"timestamp": time.Now().String(),
		}).Error("ClassicUsersDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}*/

	var classicUsers = handler.Service.FindAllPublicUsers(classicUsersDTO)
	/*if classicUsers == nil {
		fmt.Println("Ne postoji!")
		w.WriteHeader(http.StatusNotFound)
		return
	}*/

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

func (handler *ProfileSettingsHandler) FindAllNotBlockedAndMutedUsersForLoggedUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loggedUserId := vars["id"]

	var classicUsersDTO []dto.ClassicUserDTO
	if err := json.NewDecoder(r.Body).Decode(&classicUsersDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsHandler",
			"action":    "FindAllNotBlockedAndMutedUsersForLoggedUser",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to ClassicUsersDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	var profileSettingsForLoggedUser = handler.Service.FindProfileSettingByUserId(uuid.MustParse(loggedUserId))
	if profileSettingsForLoggedUser==nil{
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsHandler",
			"action":    "FindAllNotBlockedAndMutedUsersForLoggedUser",
			"timestamp": time.Now().String(),
		}).Error("Not founded profile settings for logged user!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	var usersThatLoggedUserBlocked, profilesSettingsForUsersWhoBlockedLoggedUser = handler.ProfileSettingsBlockedProfilesHandler.FindAllBlockedAndBlockingUsersForLoggedUser(profileSettingsForLoggedUser.ID, profileSettingsForLoggedUser.UserId)
	var usersForUsersWhoBlockedLoggedUser = handler.getListUserIdForListIdsOfProfilesSettings(profilesSettingsForUsersWhoBlockedLoggedUser)

	var usersThatLoggedUserMuted = handler.ProfileSettingsMutedProfilesHandler.FindAllMutedUserForLoggedUser(profileSettingsForLoggedUser.ID)
	var usersWithoutUsersThatLoggedUserBlocked = removeFromListIfExsist(classicUsersDTO, usersThatLoggedUserBlocked)
	var usersWithoutUsersThatLoggedUserBlockedAndUsersWhoBlockedLoggedUser = removeFromListIfExsist(usersWithoutUsersThatLoggedUserBlocked, usersForUsersWhoBlockedLoggedUser)

	var retVal = removeFromListIfExsist(usersWithoutUsersThatLoggedUserBlockedAndUsersWhoBlockedLoggedUser, usersThatLoggedUserMuted)

	listClassicUsersJson,_ := json.Marshal(retVal)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsHandler",
		"action":    "FindAllNotBlockedAndMutedUsersForLoggedUser",
		"timestamp": time.Now().String(),
	}).Info("Successfully found users who not blocked by logged user or not blocked logged user or logged user didn't mute!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(listClassicUsersJson)
	return
}

func removeFromListIfExsist(list []dto.ClassicUserDTO, items []uuid.UUID) []dto.ClassicUserDTO{
	var retValues []dto.ClassicUserDTO
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(list); j++ {
			if list[j].ID != items[i] && !exsistInList(retValues, items[i]){
				retValues = append(retValues, list[j])
			}
		}
	}
	return retValues
}

func exsistInList(list []dto.ClassicUserDTO, item uuid.UUID) bool{
	for j := 0; j < len(list); j++ {
		if list[j].ID == item {
			return true
		}
	}
	return false
}

func (handler *ProfileSettingsHandler) getListUserIdForListIdsOfProfilesSettings(profilesSettingsForUsersWhoBlockedLoggedUser []uuid.UUID) []uuid.UUID{
	var retVal []uuid.UUID
	for i := 0; i < len(profilesSettingsForUsersWhoBlockedLoggedUser); i++ {
		retVal = append(retVal, handler.Service.FindUserIdForProfileSettingsId(profilesSettingsForUsersWhoBlockedLoggedUser[i]))
	}
	return retVal
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
