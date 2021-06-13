package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type ClassicUserCampaignsHandler struct {
	Service * service.ClassicUserCampaignsService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *ClassicUserCampaignsHandler) CreateClassicUserCampaigns(w http.ResponseWriter, r *http.Request) {
	var classicUserCampaignsDTO dto.ClassicUserCampaignsDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserCampaignsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserCampaignsHandler",
			"action":   "CRCLASUSCAMP802",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ClassicUserCampaignsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	classicUserCampaigns := model.ClassicUserCampaigns{
		ID:            uuid.UUID{},
		ClassicUserId: classicUserCampaignsDTO.ClassicUserId,
		CampaignId:    classicUserCampaignsDTO.CampaignId,
	}

	err = handler.Service.CreateClassicUserCampaigns(&classicUserCampaigns)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserCampaignsHandler",
			"action":   "CRCLASUSCAMP802",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating campaigns for classic user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserCampaignsHandler",
		"action":   "CRCLASUSCAMP802",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created campaigns for classic user!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
