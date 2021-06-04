package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/service"
	"net/http"
	_ "strconv"
)

type CampaignChosenGroupHandler struct {
	Service *service.CampaignChosenGroupService
}

func (handler *CampaignChosenGroupHandler) CreateCampaignChosenGroup(w http.ResponseWriter, r *http.Request) {
	var campaignChosenGroupDTO dto.CampaignChosenGroupDTO
	err := json.NewDecoder(r.Body).Decode(&campaignChosenGroupDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	campaignChosenGroup := model.CampaignChosenGroup{
		ID:                     uuid.UUID{},
		CampaignId:             campaignChosenGroupDTO.CampaignId,
		RegisteredUserCategory: campaignChosenGroupDTO.RegisteredUserCategory,
	}

	err = handler.Service.CreateCampaignChosenGroup(&campaignChosenGroup)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
