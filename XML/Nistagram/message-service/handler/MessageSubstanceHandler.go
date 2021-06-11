package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/service"
	"net/http"
	"time"
)

type MessageSubstanceHandler struct {
	Service * service.MessageSubstanceService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *MessageSubstanceHandler) CreateMessageSubstance(w http.ResponseWriter, r *http.Request) {
	var messageSubstanceDTO dto.MessageSubstanceDTO
	err := json.NewDecoder(r.Body).Decode(&messageSubstanceDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageSubstanceHandler",
			"action":   "CRMESUE700",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to MessageSubstanceDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageSubstance := model.MessageSubstance{
		ID:   uuid.UUID{},
		Text: messageSubstanceDTO.Text,
	}

	err = handler.Service.CreateMessageSubstance(&messageSubstance)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageSubstanceHandler",
			"action":   "CRMESUE700",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating message substance!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "MessageSubstanceHandler",
		"action":   "CRMESUE700",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created message substance!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
