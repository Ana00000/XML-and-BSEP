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

type AdvertisementHandler struct {
	Service * service.AdvertisementService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *AdvertisementHandler) CreateAdvertisement(w http.ResponseWriter, r *http.Request) {
	var advertisementDTO dto.AdvertisementDTO
	err := json.NewDecoder(r.Body).Decode(&advertisementDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdvertisementHandler",
			"action":   "CRADA731",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to AdvertisementDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	advertisement := model.Advertisement{
		ID:                     uuid.UUID{},
		CampaignId:          advertisementDTO.CampaignId,
	}

	err = handler.Service.CreateAdvertisement(&advertisement)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdvertisementHandler",
			"action":   "CRADA731",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating advertisement!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "AdvertisementHandler",
		"action":   "CRADA731",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created advertisement!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
