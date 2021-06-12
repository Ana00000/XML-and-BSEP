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

type PostMessageSubstanceHandler struct {
	Service * service.PostMessageSubstanceService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *PostMessageSubstanceHandler) CreatePostMessageSubstance(w http.ResponseWriter, r *http.Request) {
	var postMessageSubstanceDTO dto.PostMessageSubstanceDTO
	err := json.NewDecoder(r.Body).Decode(&postMessageSubstanceDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostMessageSubstanceHandler",
			"action":   "CRPOMESUB667",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostMessageSubstanceDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postMessageSubstance := model.PostMessageSubstance{
		MessageSubstance: model.MessageSubstance{
			ID:   uuid.UUID{},
			Text: postMessageSubstanceDTO.Text,
		},
		PostId: postMessageSubstanceDTO.PostId,
	}

	err = handler.Service.CreatePostMessageSubstance(&postMessageSubstance)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostMessageSubstanceHandler",
			"action":   "CRPOMESUB667",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating post message substance!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostMessageSubstanceHandler",
		"action":   "CRPOMESUB667",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created post message substance!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
