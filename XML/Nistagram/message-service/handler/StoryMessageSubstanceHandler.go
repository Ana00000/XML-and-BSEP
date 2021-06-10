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

type StoryMessageSubstanceHandler struct {
	Service * service.StoryMessageSubstanceService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *StoryMessageSubstanceHandler) CreateStoryMessageSubstance(w http.ResponseWriter, r *http.Request) {
	var storyMessageSubstanceDTO dto.StoryMessageSubstanceDTO
	err := json.NewDecoder(r.Body).Decode(&storyMessageSubstanceDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryMessageSubstanceHandler",
			"action":   "CRSTMESUJ400",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryMessageSubstanceDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyMessageSubstance := model.StoryMessageSubstance{
		MessageSubstance: model.MessageSubstance{
			ID:   uuid.UUID{},
			Text: storyMessageSubstanceDTO.Text,
		},
		StoryId: storyMessageSubstanceDTO.StoryId,
	}

	err = handler.Service.CreateStoryMessageSubstance(&storyMessageSubstance)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryMessageSubstanceHandler",
			"action":   "CRSTMESUJ400",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating story message substance!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryMessageSubstanceHandler",
		"action":   "CRSTMESUJ400",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created story message substance!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}