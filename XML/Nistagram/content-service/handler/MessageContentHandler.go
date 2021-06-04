package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"net/http"
	_ "strconv"
)

type MessageContentHandler struct {
	Service *service.MessageContentService
}

func (handler *MessageContentHandler) CreateMessageContent(w http.ResponseWriter, r *http.Request) {
	var messageContentDTO dto.MessageContentDTO
	err := json.NewDecoder(r.Body).Decode(&messageContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch messageContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	messageContent := model.MessageContent{
		Content: model.Content{
			ID:   uuid.UUID{},
			Path: messageContentDTO.Path,
			Type: contentType,
		},
		MessageSubstanceId: messageContentDTO.MessageSubstanceId,
	}

	err = handler.Service.CreateMessageContent(&messageContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
