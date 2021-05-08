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

type ProfileSettingsMutedProfilesHandler struct {
	Service * service.ProfileSettingsMutedProfilesService
}

func (handler *ProfileSettingsMutedProfilesHandler) CreateProfileSettingsMutedProfiles(w http.ResponseWriter, r *http.Request) {
	var profileSettingsMutedProfilesDTO dto.ProfileSettingsMutedProfilesDTO
	err := json.NewDecoder(r.Body).Decode(&profileSettingsMutedProfilesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profileSettingsMutedProfiles := model.ProfileSettingsMutedProfiles{
		ID:                uuid.UUID{},
		ProfileSettingsId: profileSettingsMutedProfilesDTO.ProfileSettingsId,
		MutedProfileId:    profileSettingsMutedProfilesDTO.MutedProfileId,
	}

	err = handler.Service.CreateProfileSettingsMutedProfiles(&profileSettingsMutedProfiles)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}