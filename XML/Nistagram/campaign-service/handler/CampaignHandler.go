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

type CampaignHandler struct {
	Service * service.CampaignService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaignDTO dto.CampaignDTO
	err := json.NewDecoder(r.Body).Decode(&campaignDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CampaignHandler",
			"action":   "CRCAE175",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to CampaignDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	expTime,_ :=time.Parse(layout,campaignDTO.ExposureTime)
	campaign := model.Campaign{
		ID:                     uuid.UUID{},
		//Advertisements:         nil,
		ExposureTime:           expTime,
	}

	err = handler.Service.CreateCampaign(&campaign)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CampaignHandler",
			"action":   "CRCAE175",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating campaign!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "CampaignHandler",
		"action":   "CRCAE175",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created campaign!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
