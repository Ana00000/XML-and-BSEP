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

type StoryTagHandler struct {
	Service * service.StoryTagService
}

func (handler *StoryTagHandler) CreateStoryTag(w http.ResponseWriter, r *http.Request) {
	var storyTagDTO dto.StoryTagDTO
	err := json.NewDecoder(r.Body).Decode(&storyTagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyTag := model.StoryTag{
		Tag: model.Tag{
			ID:   uuid.UUID{},
			Name: storyTagDTO.Name,
		},
	}

	err = handler.Service.CreateStoryTag(&storyTag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}