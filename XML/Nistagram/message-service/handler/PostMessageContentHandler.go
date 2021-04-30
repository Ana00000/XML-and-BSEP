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

type PostMessageContentHandler struct {
	Service * service.PostMessageContentService
}

func (handler *PostMessageContentHandler) CreatePostMessageContent(w http.ResponseWriter, r *http.Request) {
	var postMessageContentDTO dto.PostMessageContentDTO
	err := json.NewDecoder(r.Body).Decode(&postMessageContentDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postMessageContent := model.PostMessageContent{
		MessageContent : model.MessageContent{
			ID: uuid.UUID{},
			Text: postMessageContentDTO.Text,
		},
		// Post
	}

	err = handler.Service.CreatePostMessageContent(&postMessageContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}