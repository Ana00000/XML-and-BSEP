package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type MessageHandler struct {
	Service * service.MessageService
}

func (handler *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var messageDTO dto.MessageDTO
	err := json.NewDecoder(r.Body).Decode(&messageDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate, _ := time.Parse(layout, messageDTO.CreationDate)

	message := model.Message{
		ID: uuid.UUID{},
		MessageSubstanceId: messageDTO.MessageContentID,
		IsDisposable: messageDTO.IsDisposable,
		CreationDate: creationDate,
		SenderUserID: messageDTO.SenderUserID,
        ReceiverUserID: messageDTO.ReceiverUserID,
        IsDeleted: messageDTO.IsDeleted,
	}

	err = handler.Service.CreateMessage(&message)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}