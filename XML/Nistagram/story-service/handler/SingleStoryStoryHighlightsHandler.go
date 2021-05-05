package handler

import (
	"../dto"
	"../model"
	"../service"
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
