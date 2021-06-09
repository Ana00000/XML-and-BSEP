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

type ProfileSettingsRejectedMessageProfilesHandler struct {
	Service   *service.ProfileSettingsRejectedMessageProfilesService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *ProfileSettingsRejectedMessageProfilesHandler) CreateProfileSettingsRejectedMessageProfiles(w http.ResponseWriter, r *http.Request) {
	var profileSettingsRejectedMessageProfilesDTO dto.ProfileSettingsRejectedMessageProfilesDTO
	if err := json.NewDecoder(r.Body).Decode(&profileSettingsRejectedMessageProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsRejectedMessageProfilesHandler",
			"action":    "CRPROFSETTINGSREJCTMESSPROF1802",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to ProfileSettingsRejectedMessageProfilesDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&profileSettingsRejectedMessageProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsRejectedMessageProfilesHandler",
			"action":    "CRPROFSETTINGSREJCTMESSPROF1802",
			"timestamp": time.Now().String(),
		}).Error("ProfileSettingsRejectedMessageProfilesDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	profileSettingsRejectedMessageProfiles := model.ProfileSettingsRejectedMessageProfiles{
		ID:                       uuid.UUID{},
		ProfileSettingsId:        profileSettingsRejectedMessageProfilesDTO.ProfileSettingsId,
		RejectedMessageProfileId: profileSettingsRejectedMessageProfilesDTO.RejectedMessageProfileId,
	}

	if err := handler.Service.CreateProfileSettingsRejectedMessageProfiles(&profileSettingsRejectedMessageProfiles); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsRejectedMessageProfilesHandler",
			"action":    "CRPROFSETTINGSREJCTMESSPROF1802",
			"timestamp": time.Now().String(),
		}).Error("Failed creating profile settings rejected message profiles!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsRejectedMessageProfilesHandler",
		"action":    "CRPROFSETTINGSREJCTMESSPROF1802",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings rejected message profiles!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
