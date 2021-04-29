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

type ProfileSettingsHandler struct {
	Service * service.ProfileSettingsService
}

func (handler *ProfileSettingsHandler) CreateProfileSettings(w http.ResponseWriter, r *http.Request) {
	var profileSettingsDTO dto.ProfileSettingsDTO
	err := json.NewDecoder(r.Body).Decode(&profileSettingsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profileSettings := model.ProfileSettings{
		ID:          uuid.UUID{},
		UserId: profileSettingsDTO.UserId,
		UserVisibility:      profileSettingsDTO.UserVisibility,
		MessageApprovalType:       profileSettingsDTO.MessageApprovalType,
	}

	err = handler.Service.CreateProfileSettings(&profileSettings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

