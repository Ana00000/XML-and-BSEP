package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type StoryTagStoriesHandler struct {
	Service * service.StoryTagStoriesService
}

func (handler *StoryTagStoriesHandler) CreateStoryTagStories(w http.ResponseWriter, r *http.Request) {
	var storyTagStoriesDTO dto.StoryTagStoriesDTO
	err := json.NewDecoder(r.Body).Decode(&storyTagStoriesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyTagStories := model.StoryTagStories{
		ID:         uuid.UUID{},
		StoryTagId: storyTagStoriesDTO.StoryTagId,
		StoryId:    storyTagStoriesDTO.StoryId,
	}

	err = handler.Service.CreateStoryTagStories(&storyTagStories)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}