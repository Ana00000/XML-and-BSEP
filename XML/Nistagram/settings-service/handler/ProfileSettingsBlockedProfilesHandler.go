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
	_ "strconv"
	"time"
)

type ProfileSettingsBlockedProfilesHandler struct {
	Service   *service.ProfileSettingsBlockedProfilesService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *ProfileSettingsBlockedProfilesHandler) CreateProfileSettingsBlockedProfiles(w http.ResponseWriter, r *http.Request) {
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
