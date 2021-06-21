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

type ProfileSettingsMutedProfilesHandler struct {
	Service  *service.ProfileSettingsMutedProfilesService
	ProfileSettingsService * service.ProfileSettingsService
	LogInfo  *logrus.Logger
	LogError *logrus.Logger
	Validator *validator.Validate
}

func (handler *ProfileSettingsMutedProfilesHandler) CreateProfileSettingsMutedProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var profileSettingsMutedProfilesDTO dto.ProfileSettingsMutedProfilesDTO

	if err := json.NewDecoder(r.Body).Decode(&profileSettingsMutedProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ProfileSettingsMutedProfilesHandler",
			"action": "CRPROFSETTINGSMUTPROF7777",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to ProfileSettingsMutedProfilesDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&profileSettingsMutedProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsMutedProfilesHandler",
			"action":    "CRPROFSETTINGSMUTPROF7777",
			"timestamp": time.Now().String(),
		}).Error("ProfileSettingsMutedProfilesDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	profileSettingsMutedProfiles := model.ProfileSettingsMutedProfiles{
		ID:                uuid.UUID{},
		ProfileSettingsId: profileSettingsMutedProfilesDTO.ProfileSettingsId,
		MutedProfileId:    profileSettingsMutedProfilesDTO.MutedProfileId,
	}

	if err := handler.Service.CreateProfileSettingsMutedProfiles(&profileSettingsMutedProfiles); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ProfileSettingsMutedProfilesHandler",
			"action":   "CRPROFSETTINGSMUTPROF7777",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating profile settings muted profiles!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsMutedProfilesHandler",
		"action":    "CRPROFSETTINGSMUTPROF7777",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings muted profiles!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ProfileSettingsMutedProfilesHandler) MuteUser(w http.ResponseWriter, r *http.Request) {

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ProfileSettingsMutedProfilesHandler",
			"action":   "MuteUser",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-mute-user-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ProfileSettingsMutedProfilesHandler",
			"action":   "MuteUser",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}

	var muteUserDTO dto.MuteUserDTO
	if err := json.NewDecoder(r.Body).Decode(&muteUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsMutedProfilesHandler",
			"action":    "MuteUser",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to MuteUserDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var profileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(muteUserDTO.LoggedInUser)

	var profileSettingsMutedProfiles = model.ProfileSettingsMutedProfiles{
		ID:                uuid.UUID{},
		ProfileSettingsId: profileSettings.ID,
		MutedProfileId:    muteUserDTO.MutedUser,
	}

	if err := handler.Service.CreateProfileSettingsMutedProfiles(&profileSettingsMutedProfiles); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsMutedProfilesHandler",
			"action":    "MuteUser",
			"timestamp": time.Now().String(),
		}).Error("Failed creating profile settings muted profiles!")
		//fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsMutedProfilesHandler",
		"action":    "MuteUser",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings muted profiles!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ProfileSettingsMutedProfilesHandler) UnmuteUser(w http.ResponseWriter, r *http.Request) {
	var muteUserDTO dto.MuteUserDTO
	if err := json.NewDecoder(r.Body).Decode(&muteUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsMutedProfilesHandler",
			"action":    "UnmuteUser",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to MuteUserDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var profileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(muteUserDTO.LoggedInUser)

	var profileSettingsMutedProfiles = handler.Service.FindProfileSettingsMutedProfiles(profileSettings.ID,muteUserDTO.MutedUser)

	if profileSettingsMutedProfiles==nil{
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsMutedProfilesHandler",
			"action":    "UnmuteUser",
			"timestamp": time.Now().String(),
		}).Error("User isn't muted")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	handler.Service.UnmuteUser(profileSettingsMutedProfiles.ID)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
