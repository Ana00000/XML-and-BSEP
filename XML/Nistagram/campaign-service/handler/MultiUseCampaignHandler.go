package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
	"time"
)

type MultiUseCampaignHandler struct {
	Service * service.MultiUseCampaignService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *MultiUseCampaignHandler) CreateMultiUseCampaign(w http.ResponseWriter, r *http.Request) {
	var multiUseCampaignDTO dto.MultiUseCampaignDTO
	err := json.NewDecoder(r.Body).Decode(&multiUseCampaignDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MultiUseCampaignHandler",
			"action":   "CRMUUSCAP780",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to MultiUseCampaignDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	exposureTime,_ :=time.Parse(layout,multiUseCampaignDTO.ExposureTime)
	expiryTime,_ :=time.Parse(layout,multiUseCampaignDTO.ExpiryTime)
	multiUseCampaign := model.MultiUseCampaign{
		Campaign:   model.Campaign{
			ID:                     uuid.UUID{},
			//Advertisements:         nil,
			ExposureTime:           exposureTime,
		},
		ExpiryTime: expiryTime,
		Frequency:  multiUseCampaignDTO.Frequency,
	}

	err = handler.Service.CreateMultiUseCampaign(&multiUseCampaign)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MultiUseCampaignHandler",
			"action":   "CRMUUSCAP780",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating multi use campaign!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "MultiUseCampaignHandler",
		"action":   "CRMUUSCAP780",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created multi use campaign!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
