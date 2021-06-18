package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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

type ProfileSettingsBlockedProfilesHandler struct {
	Service   *service.ProfileSettingsBlockedProfilesService
	ProfileSettingsService *service.ProfileSettingsService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *ProfileSettingsBlockedProfilesHandler) CreateProfileSettingsBlockedProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var profileSettingsBlockedProfilesDTO dto.ProfileSettingsBlockedProfilesDTO
	if err := json.NewDecoder(r.Body).Decode(&profileSettingsBlockedProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsBlockedProfilesHandler",
			"action":    "CRPROFSETTINGSBLOCPROF101",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to ProfileSettingsBlockedProfilesDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&profileSettingsBlockedProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsBlockedProfilesHandler",
			"action":    "CRPROFSETTINGSBLOCPROF101",
			"timestamp": time.Now().String(),
		}).Error("ProfileSettingsBlockedProfilesDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	profileSettingsBlockedProfiles := model.ProfileSettingsBlockedProfiles{
		ID:                uuid.UUID{},
		ProfileSettingsId: profileSettingsBlockedProfilesDTO.ProfileSettingsId,
		BlockedProfileId:  profileSettingsBlockedProfilesDTO.BlockedProfileId,
	}

	if err := handler.Service.CreateProfileSettingsBlockedProfiles(&profileSettingsBlockedProfiles); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsBlockedProfilesHandler",
			"action":    "CRPROFSETTINGSBLOCPROF101",
			"timestamp": time.Now().String(),
		}).Error("Failed creating profile settings blocked profiles!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsBlockedProfilesHandler",
		"action":    "CRPROFSETTINGSBLOCPROF101",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings blocked profiles!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ProfileSettingsBlockedProfilesHandler) BlockUser(w http.ResponseWriter, r *http.Request) {

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ProfileSettingsBlockedProfilesHandler",
			"action":   "BlockUser",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-block-user-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ProfileSettingsBlockedProfilesHandler",
			"action":   "BlockUser",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}

	var blockUserDTO dto.BlockUserDTO
	if err := json.NewDecoder(r.Body).Decode(&blockUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsBlockedProfilesHandler",
			"action":    "BlockUser",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to BlockUserDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var profileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(blockUserDTO.LoggedInUser)

	var profileSettingsBlockedProfiles = model.ProfileSettingsBlockedProfiles{
		ID:                uuid.UUID{},
		ProfileSettingsId: profileSettings.ID,
		BlockedProfileId:  blockUserDTO.BlockedUser,
	}

	if err := handler.Service.CreateProfileSettingsBlockedProfiles(&profileSettingsBlockedProfiles); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsBlockedProfilesHandler",
			"action":    "BlockUser",
			"timestamp": time.Now().String(),
		}).Error("Failed creating profile settings blocked profiles!")
		//fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsBlockedProfilesHandler",
		"action":    "BlockUser",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings blocked profiles!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

}
