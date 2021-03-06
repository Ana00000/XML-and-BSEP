package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type DisposableCampaignHandler struct {
	Service * service.DisposableCampaignService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *DisposableCampaignHandler) CreateDisposableCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var disposableCampaignDTO dto.DisposableCampaignDTO
	err := json.NewDecoder(r.Body).Decode(&disposableCampaignDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "DisposableCampaignHandler",
			"action":   "CRDICAF756",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to DisposableCampaignDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	expTime, _ := time.Parse(layout, disposableCampaignDTO.ExposureTime)
	disposableCampaign := model.DisposableCampaign{
		Campaign: model.Campaign{
			ID: uuid.UUID{},
			ExposureTime: expTime,
		},
	}

	err = handler.Service.CreateDisposableCampaign(&disposableCampaign)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "DisposableCampaignHandler",
			"action":   "CRDICAF756",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating disposable campaign!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "DisposableCampaignHandler",
		"action":   "CRDICAF756",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created disposable campaign!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
