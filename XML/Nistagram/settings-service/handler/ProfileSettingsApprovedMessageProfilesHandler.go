package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type ProfileSettingsApprovedMessageProfilesHandler struct {
	Service * service.ProfileSettingsApprovedMessageProfilesService
}

func (handler *ProfileSettingsApprovedMessageProfilesHandler) CreateProfileSettingsApprovedMessageProfiles(w http.ResponseWriter, r *http.Request) {
	var profileSettingsApprovedMessageProfilesDTO dto.ProfileSettingsApprovedMessageProfilesDTO
	err := json.NewDecoder(r.Body).Decode(&profileSettingsApprovedMessageProfilesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profileSettingsApprovedMessageProfiles := model.ProfileSettingsApprovedMessageProfiles{
		ID:                       uuid.UUID{},
		ProfileSettingsId:        profileSettingsApprovedMessageProfilesDTO.ProfileSettingsId,
		ApprovedMessageProfileId: profileSettingsApprovedMessageProfilesDTO.ApprovedMessageProfileId,
	}

	err = handler.Service.CreateProfileSettingsApprovedMessageProfiles(&profileSettingsApprovedMessageProfiles)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
