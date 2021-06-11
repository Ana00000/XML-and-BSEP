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

type ProfileSettingsMutedProfilesHandler struct {
	Service  *service.ProfileSettingsMutedProfilesService
	LogInfo  *logrus.Logger
	LogError *logrus.Logger
	Validator *validator.Validate
}

func (handler *ProfileSettingsMutedProfilesHandler) CreateProfileSettingsMutedProfiles(w http.ResponseWriter, r *http.Request) {
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
