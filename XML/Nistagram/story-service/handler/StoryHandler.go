package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type StoryHandler struct {
	Service * service.StoryService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}
//CRSTRY90211
func (handler *StoryHandler) CreateStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var storyDTO dto.StoryDTO
	err := json.NewDecoder(r.Body).Decode(&storyDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryHandler",
			"action":   "CRSTRY90211",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to StoryDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyType := model.CLOSE_FRIENDS
	switch storyDTO.Type {
	case "ALL_FRIENDS":
		storyType = model.ALL_FRIENDS
	case "PUBLIC":
		storyType = model.PUBLIC
	}

	id := uuid.New()
	story := model.Story{
		ID:           id,
		CreationDate: time.Now(),
		Description:  storyDTO.Description,
		UserId:       storyDTO.UserId,
		LocationId:   storyDTO.LocationId,
		IsDeleted:    false,
		Type:         storyType,
	}

	err = handler.Service.CreateStory(&story)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "StoryHandler",
			"action":   "CRSTRY90211",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating story!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "StoryHandler",
		"action":   "CRSTRY90211",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created story!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
