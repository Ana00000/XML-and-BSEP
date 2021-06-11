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

type ProfileSettingsApprovedMessageProfilesHandler struct {
	Service   *service.ProfileSettingsApprovedMessageProfilesService
	LogInfo   *logrus.Logger
	LogError  *logrus.Logger
	Validator *validator.Validate
}

func (handler *ProfileSettingsApprovedMessageProfilesHandler) CreateProfileSettingsApprovedMessageProfiles(w http.ResponseWriter, r *http.Request) {
	var profileSettingsApprovedMessageProfilesDTO dto.ProfileSettingsApprovedMessageProfilesDTO

	if err := json.NewDecoder(r.Body).Decode(&profileSettingsApprovedMessageProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsApprovedMessageProfilesHandler",
			"action":    "CRPROFSETTINGSAPPROVEDMESPROF967",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast jason to ProfileSettingsApprovedMessageProfilesDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&profileSettingsApprovedMessageProfilesDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsApprovedMessageProfilesHandler",
			"action":    "CRPROFSETTINGSAPPROVEDMESPROF967",
			"timestamp": time.Now().String(),
		}).Error("ProfileSettingsApprovedMessageProfilesDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	profileSettingsApprovedMessageProfiles := model.ProfileSettingsApprovedMessageProfiles{
		ID:                       uuid.UUID{},
		ProfileSettingsId:        profileSettingsApprovedMessageProfilesDTO.ProfileSettingsId,
		ApprovedMessageProfileId: profileSettingsApprovedMessageProfilesDTO.ApprovedMessageProfileId,
	}

	if err := handler.Service.CreateProfileSettingsApprovedMessageProfiles(&profileSettingsApprovedMessageProfiles); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ProfileSettingsApprovedMessageProfilesHandler",
			"action":    "CRPROFSETTINGSAPPROVEDMESPROF967",
			"timestamp": time.Now().String(),
		}).Error("Failed creating profile settings approved message profiles!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ProfileSettingsApprovedMessageProfilesHandler",
		"action":    "CRPROFSETTINGSAPPROVEDMESPROF967",
		"timestamp": time.Now().String(),
	}).Info("Successfully created profile settings approved message profiles!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
