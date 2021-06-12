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

type MessageHandler struct {
	Service * service.MessageService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var messageDTO dto.MessageDTO
	err := json.NewDecoder(r.Body).Decode(&messageDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageHandler",
			"action":   "CRMEE454",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to MessageDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate, _ := time.Parse(layout, messageDTO.CreationDate)

	message := model.Message{
		ID:                 uuid.UUID{},
		MessageSubstanceId: messageDTO.MessageContentID,
		IsDisposable:       messageDTO.IsDisposable,
		CreationDate:       creationDate,
		SenderUserID:       messageDTO.SenderUserID,
		ReceiverUserID:     messageDTO.ReceiverUserID,
		IsDeleted:          messageDTO.IsDeleted,
	}

	err = handler.Service.CreateMessage(&message)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageHandler",
			"action":   "CRMEE454",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating message!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "MessageHandler",
		"action":   "CRMEE454",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created message!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
