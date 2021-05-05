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

type StoryHighlightHandler struct {
	Service * service.StoryHighlightService
}

func (handler *StoryHighlightHandler) CreateStoryHighlight(w http.ResponseWriter, r *http.Request) {
	var storyHighlightDTO dto.StoryHighlightDTO
	err := json.NewDecoder(r.Body).Decode(&storyHighlightDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyHighlight := model.StoryHighlight{
		ID:          uuid.UUID{},
		Title: storyHighlightDTO.Title,
		UserId:      storyHighlightDTO.UserId,
		//Stories: nil,
	}

	err = handler.Service.CreateStoryHighlight(&storyHighlight)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
