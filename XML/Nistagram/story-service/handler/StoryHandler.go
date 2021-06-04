package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type StoryHandler struct {
	Service *service.StoryService
}

func (handler *StoryHandler) CreateStory(w http.ResponseWriter, r *http.Request) {
	var storyDTO dto.StoryDTO
	err := json.NewDecoder(r.Body).Decode(&storyDTO)
	if err != nil {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
