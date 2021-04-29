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

type MessageContentHandler struct {
	Service * service.MessageContentService
}

func (handler *MessageContentHandler) CreateMessageContent(w http.ResponseWriter, r *http.Request) {
	var messageContentDTO dto.MessageContentDTO
	err := json.NewDecoder(r.Body).Decode(&messageContentDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageContent := model.MessageContent{
		ID: uuid.UUID{},
		Text: messageContentDTO.Text,
	}

	err = handler.Service.CreateMessageContent(&messageContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}