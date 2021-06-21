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

type ProfileSettingsStoryNotificationsProfilesHandler struct {
	Service   *service.ProfileSettingsStoryNotificationsProfilesService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *ProfileSettingsStoryNotificationsProfilesHandler) CreateProfileSettingsStoryNotificationsProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var profileSettingsStoryNotificationsProfilesDTO dto.ProfileSettingsStoryNotificationsProfilesDTO
	if err := json.NewDecoder(r.Body).Decode(&profileSettingsStoryNotificationsProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsStoryNotificationsProfilesHandler",
			"action":    "CreateProfileSettingsStoryNotificationsProfiles",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to ProfileSettingsStoryNotificationsProfilesDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&profileSettingsStoryNotificationsProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsStoryNotificationsProfilesHandler",
			"action":    "CreateProfileSettingsStoryNotificationsProfiles",
			"timestamp": time.Now().String(),
		}).Error("ProfileSettingsStoryNotificationsProfilesDTO fields aren not in the valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	profileSettingsStoryNotificationsProfiles := model.ProfileSettingsStoryNotificationsProfiles{
		ID:                uuid.UUID{},
		ProfileSettingsId: profileSettingsStoryNotificationsProfilesDTO.ProfileSettingsId,
		StoryNotificationsProfileId:  profileSettingsStoryNotificationsProfilesDTO.StoryNotificationsProfileId,
	}

	if err := handler.Service.CreateProfileSettingsStoryNotificationsProfiles(&profileSettingsStoryNotificationsProfiles); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsStoryNotificationsProfilesHandler",
			"action":    "CreateProfileSettingsStoryNotificationsProfiles",
			"timestamp": time.Now().String(),
		}).Error("Failed creating profile settings story notifications profiles!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsStoryNotificationsProfilesHandler",
		"action":    "CreateProfileSettingsStoryNotificationsProfiles",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings story notifications profiles!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

