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

type StoryHandler struct {
	Service * service.StoryService
}

func (handler *StoryHandler) CreateStory(w http.ResponseWriter, r *http.Request) {
	var storyDTO dto.StoryDTO
	err := json.NewDecoder(r.Body).Decode(&storyDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	story := model.Story{
		ID:          uuid.UUID{},
		CreationDate: storyDTO.CreationDate,
		UserId:      storyDTO.UserId,
		Location:      storyDTO.Location,
		IsDeleted:      storyDTO.IsDeleted,
		Type:      storyDTO.Type,

	}

	err = handler.Service.CreateStory(&story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

