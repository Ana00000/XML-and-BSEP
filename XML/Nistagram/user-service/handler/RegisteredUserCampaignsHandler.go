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

type RegisteredUserCampaignsHandler struct {
	Service * service.RegisteredUserCampaignsService
}

func (handler *RegisteredUserCampaignsHandler) CreateRegisteredUserCampaigns(w http.ResponseWriter, r *http.Request) {
	var registeredUserCampaignsDTO dto.RegisteredUserCampaignsDTO
	err := json.NewDecoder(r.Body).Decode(&registeredUserCampaignsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	registeredUserCampaigns := model.RegisteredUserCampaigns{
		ID:               uuid.UUID{},
		RegisteredUserId: registeredUserCampaignsDTO.RegisteredUserId,
		CampaignId:       registeredUserCampaignsDTO.CampaignId,
	}

	err = handler.Service.CreateRegisteredUserCampaigns(&registeredUserCampaigns)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}