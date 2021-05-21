package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type SingleStoryStoryHighlightsHandler struct {
	Service * service.SingleStoryStoryHighlightsService
}

func (handler *SingleStoryStoryHighlightsHandler) CreateSingleStoryStoryHighlights(w http.ResponseWriter, r *http.Request) {
	var singleStoryStoryHighlightsDTO dto.SingleStoryStoryHighlightsDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryStoryHighlightsDTO)
	if err != nil {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SingleStoryStoryHighlightsHandler) FindAllSingleStoryStoryHighlightsForStory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	singleStoryStoryHighlights := handler.Service.FindAllSingleStoryStoryHighlightsForStory(uuid.MustParse(id))
	singleStoryStoryHighlightsJson, _ := json.Marshal(singleStoryStoryHighlights)
	if singleStoryStoryHighlightsJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(singleStoryStoryHighlightsJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}