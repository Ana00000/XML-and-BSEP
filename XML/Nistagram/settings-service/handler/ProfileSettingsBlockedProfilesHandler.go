package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type ProfileSettingsBlockedProfilesHandler struct {
	Service * service.ProfileSettingsBlockedProfilesService
}

func (handler *ProfileSettingsBlockedProfilesHandler) CreateProfileSettingsBlockedProfiles(w http.ResponseWriter, r *http.Request) {
	var profileSettingsBlockedProfilesDTO dto.ProfileSettingsBlockedProfilesDTO
	err := json.NewDecoder(r.Body).Decode(&profileSettingsBlockedProfilesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profileSettingsBlockedProfiles := model.ProfileSettingsBlockedProfiles{
		ID:                uuid.UUID{},
		ProfileSettingsId: profileSettingsBlockedProfilesDTO.ProfileSettingsId,
		BlockedProfileId:  profileSettingsBlockedProfilesDTO.BlockedProfileId,
	}

	err = handler.Service.CreateProfileSettingsBlockedProfiles(&profileSettingsBlockedProfiles)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
