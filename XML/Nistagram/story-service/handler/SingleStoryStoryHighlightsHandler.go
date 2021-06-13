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

type SingleStoryStoryHighlightsHandler struct {
	Service * service.SingleStoryStoryHighlightsService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//CRSINGLSTRYSTRYHIGHLHTS9820
func (handler *SingleStoryStoryHighlightsHandler) CreateSingleStoryStoryHighlights(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var singleStoryStoryHighlightsDTO dto.SingleStoryStoryHighlightsDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryStoryHighlightsDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryStoryHighlightsHandler",
			"action":   "CRSINGLSTRYSTRYHIGHLHTS9820",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to SingleStoryStoryHighlightsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	singleStoryStoryHighlights := model.SingleStoryStoryHighlights{
		ID:               uuid.UUID{},
		SingleStoryId:    singleStoryStoryHighlightsDTO.SingleStoryId,
		StoryHighlightId: singleStoryStoryHighlightsDTO.StoryHighlightId,
	}

	err = handler.Service.CreateSingleStoryStoryHighlights(&singleStoryStoryHighlights)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "SingleStoryStoryHighlightsHandler",
			"action":   "CRSINGLSTRYSTRYHIGHLHTS9820",
			"timestamp":   time.Now().String(),
		}).Error("Failed adding single story to story highlights!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "SingleStoryStoryHighlightsHandler",
		"action":   "CRSINGLSTRYSTRYHIGHLHTS9820",
		"timestamp":   time.Now().String(),
	}).Info("Successfully added single story to story highlights!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALSINGLSTRYSTRYHIGHLHTSFORSTRY9840
func (handler *SingleStoryStoryHighlightsHandler) FindAllSingleStoryStoryHighlightsForStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	singleStoryStoryHighlights := handler.Service.FindAllSingleStoryStoryHighlightsForStory(uuid.MustParse(id))
	singleStoryStoryHighlightsJson, _ := json.Marshal(singleStoryStoryHighlights)
	if singleStoryStoryHighlightsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "SingleStoryStoryHighlightsHandler",
			"action":   "FIDALSINGLSTRYSTRYHIGHLHTSFORSTRY9840",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded all single story of story highlights for story!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(singleStoryStoryHighlightsJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "SingleStoryStoryHighlightsHandler",
		"action":   "FIDALSINGLSTRYSTRYHIGHLHTSFORSTRY9840",
		"timestamp":   time.Now().String(),
	}).Error("Failed finding all single story of story highlights for story!")
	w.WriteHeader(http.StatusBadRequest)
}

//FIDALSINGLSTRYSTRYHIGHLHTSFORSTRYLHT9193
func (handler *SingleStoryStoryHighlightsHandler) FindAllSingleStoryStoryHighlightsForStoryHighlight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	singleStoryStoryHighlights := handler.Service.FindAllSingleStoryStoryHighlightsForStoryHighlight(uuid.MustParse(id))
	singleStoryStoryHighlightsJson, _ := json.Marshal(singleStoryStoryHighlights)
	if singleStoryStoryHighlightsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "SingleStoryStoryHighlightsHandler",
			"action":   "FIDALSINGLSTRYSTRYHIGHLHTSFORSTRYLHT9193",
			"timestamp":   time.Now().String(),
		}).Info("Successfully founded all single story of story highlights for story highlight!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(singleStoryStoryHighlightsJson)
		return
	}

	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "SingleStoryStoryHighlightsHandler",
		"action":   "FIDALSINGLSTRYSTRYHIGHLHTSFORSTRYLHT9193",
		"timestamp":   time.Now().String(),
	}).Error("Failed finding all single story of story highlights for story highlight!")
	w.WriteHeader(http.StatusBadRequest)
}
