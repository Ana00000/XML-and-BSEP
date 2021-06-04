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
	"time"
)

type MultiUseCampaignHandler struct {
	Service *service.MultiUseCampaignService
}

func (handler *MultiUseCampaignHandler) CreateMultiUseCampaign(w http.ResponseWriter, r *http.Request) {
	var multiUseCampaignDTO dto.MultiUseCampaignDTO
	err := json.NewDecoder(r.Body).Decode(&multiUseCampaignDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	exposureTime, _ := time.Parse(layout, multiUseCampaignDTO.ExposureTime)
	expiryTime, _ := time.Parse(layout, multiUseCampaignDTO.ExpiryTime)
	multiUseCampaign := model.MultiUseCampaign{
		Campaign: model.Campaign{
			ID: uuid.UUID{},
			ExposureTime: exposureTime,
		},
		ExpiryTime: expiryTime,
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
