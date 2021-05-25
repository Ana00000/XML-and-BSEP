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

type ProfileSettingsRejectedMessageProfilesHandler struct {
	Service * service.ProfileSettingsRejectedMessageProfilesService
}

func (handler *ProfileSettingsRejectedMessageProfilesHandler) CreateProfileSettingsRejectedMessageProfiles(w http.ResponseWriter, r *http.Request) {
	var profileSettingsRejectedMessageProfilesDTO dto.ProfileSettingsRejectedMessageProfilesDTO
	err := json.NewDecoder(r.Body).Decode(&profileSettingsRejectedMessageProfilesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profileSettingsRejectedMessageProfiles := model.ProfileSettingsRejectedMessageProfiles{
		ID:                       uuid.UUID{},
		ProfileSettingsId:        profileSettingsRejectedMessageProfilesDTO.ProfileSettingsId,
		RejectedMessageProfileId: profileSettingsRejectedMessageProfilesDTO.RejectedMessageProfileId,
	}

	err = handler.Service.CreateProfileSettingsRejectedMessageProfiles(&profileSettingsRejectedMessageProfiles)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
