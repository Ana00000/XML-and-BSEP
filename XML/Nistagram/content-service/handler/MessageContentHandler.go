package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type MessageContentHandler struct {
	Service * service.MessageContentService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *MessageContentHandler) CreateMessageContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var messageContentDTO dto.MessageContentDTO
	err := json.NewDecoder(r.Body).Decode(&messageContentDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageContentHandler",
			"action":   "CRMECOK313",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to MessageContentDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch messageContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	messageContent := model.MessageContent{
		Content: model.Content{
			ID:   uuid.UUID{},
			Path: messageContentDTO.Path,
			Type: contentType,
		},
		MessageSubstanceId: messageContentDTO.MessageSubstanceId,
	}

	err = handler.Service.CreateMessageContent(&messageContent)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "MessageContentHandler",
			"action":   "CRMECOK313",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating message content!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "MessageContentHandler",
		"action":   "CRMECOK313",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created message content!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
