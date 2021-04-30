package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type StoryMessageContentHandler struct {
	Service * service.StoryMessageContentService
}

func (handler *StoryMessageContentHandler) CreateStoryMessageContent(w http.ResponseWriter, r *http.Request) {
	var storyMessageContentDTO dto.StoryMessageContentDTO
	err := json.NewDecoder(r.Body).Decode(&storyMessageContentDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyMessageContent := model.StoryMessageContent{
		MessageContent : model.MessageContent{
			ID: uuid.UUID{},
			Text: storyMessageContentDTO.Text,
		},
		// Story
	}

	err = handler.Service.CreateStoryMessageContent(&storyMessageContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}