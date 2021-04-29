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

type MultiUseCampaignHandler struct {
	Service * service.MultiUseCampaignService
}

func (handler *MultiUseCampaignHandler) CreateMultiUseCampaign(w http.ResponseWriter, r *http.Request) {
	var multiUseCampaignDTO dto.MultiUseCampaignDTO
	err := json.NewDecoder(r.Body).Decode(&multiUseCampaignDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	expoTime,_ :=time.Parse(layout,multiUseCampaignDTO.ExposureTime)
	expiTime,_ :=time.Parse(layout,multiUseCampaignDTO.ExpiryTime)
	multiUseCampaign := model.MultiUseCampaign{
		Campaign:   model.Campaign{
			ID:                     uuid.UUID{},
			Advertisements:         nil,
			ExposureTime:           expoTime,
		},
		ExpiryTime: expiTime,
		Frequency:  multiUseCampaignDTO.Frequency,
	}

	err = handler.Service.CreateMultiUseCampaign(&multiUseCampaign)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
