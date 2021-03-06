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

type CampaignChosenGroupHandler struct {
	Service * service.CampaignChosenGroupService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *CampaignChosenGroupHandler) CreateCampaignChosenGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var campaignChosenGroupDTO dto.CampaignChosenGroupDTO
	err := json.NewDecoder(r.Body).Decode(&campaignChosenGroupDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CampaignChosenGroupHandler",
			"action":   "CRCACHGRB114",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to CampaignChosenGroupDTO!")
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
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CampaignChosenGroupHandler",
			"action":   "CRCACHGRB114",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating campaign chosen group!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "CampaignChosenGroupHandler",
		"action":   "CRCACHGRB114",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created campaign chosen group!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
