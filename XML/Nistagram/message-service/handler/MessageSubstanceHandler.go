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

type MessageSubstanceHandler struct {
	Service * service.MessageSubstanceService
}

func (handler *MessageSubstanceHandler) CreateMessageSubstance(w http.ResponseWriter, r *http.Request) {
	var messageSubstanceDTO dto.MessageSubstanceDTO
	err := json.NewDecoder(r.Body).Decode(&messageSubstanceDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageSubstance := model.MessageSubstance{
		ID: uuid.UUID{},
		Text: messageSubstanceDTO.Text,
	}

	err = handler.Service.CreateMessageSubstance(&messageSubstance)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}