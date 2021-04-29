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
	"time"
)

type CampaignHandler struct {
	Service * service.CampaignService
}

func (handler *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaignDTO dto.CampaignDTO
	err := json.NewDecoder(r.Body).Decode(&campaignDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	expTime,_ :=time.Parse(layout,campaignDTO.ExposureTime)
	campaign := model.Campaign{
		ID:                     uuid.UUID{},
		Advertisements:         nil,
		ExposureTime:           expTime,
	}

	err = handler.Service.CreateCampaign(&campaign)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
