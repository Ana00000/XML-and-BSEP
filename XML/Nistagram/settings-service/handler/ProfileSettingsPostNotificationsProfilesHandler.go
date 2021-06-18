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

type ProfileSettingsPostNotificationsProfilesHandler struct {
	Service   *service.ProfileSettingsPostNotificationsProfilesService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}



func (handler *ProfileSettingsPostNotificationsProfilesHandler) CreateProfileSettingsPostNotificationsProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var profileSettingsPostNotificationsProfilesDTO dto.ProfileSettingsPostNotificationsProfilesDTO
	if err := json.NewDecoder(r.Body).Decode(&profileSettingsPostNotificationsProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsPostNotificationsProfilesHandler",
			"action":    "CreateProfileSettingsPostNotificationsProfiles",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to ProfileSettingsPostNotificationsProfilesDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&profileSettingsPostNotificationsProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsPostNotificationsProfilesHandler",
			"action":    "CreateProfileSettingsPostNotificationsProfiles",
			"timestamp": time.Now().String(),
		}).Error("ProfileSettingsPostNotificationsProfilesDTO fields aren not in the valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	profileSettingsPostNotificationsProfiles := model.ProfileSettingsPostNotificationsProfiles{
		ID:                uuid.UUID{},
		ProfileSettingsId: profileSettingsPostNotificationsProfilesDTO.ProfileSettingsId,
		PostNotificationsProfileId:  profileSettingsPostNotificationsProfilesDTO.PostNotificationsProfileId,
	}

	if err := handler.Service.CreateProfileSettingsPostNotificationsProfiles(&profileSettingsPostNotificationsProfiles); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsPostNotificationsProfilesHandler",
			"action":    "CreateProfileSettingsPostNotificationsProfiles",
			"timestamp": time.Now().String(),
		}).Error("Failed creating profile settings post notifications profiles!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsPostNotificationsProfilesHandler",
		"action":    "CreateProfileSettingsPostNotificationsProfiles",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings post notifications profiles!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}




