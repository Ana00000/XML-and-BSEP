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

type PostMessageSubstanceHandler struct {
	Service * service.PostMessageSubstanceService
}

func (handler *PostMessageSubstanceHandler) CreatePostMessageSubstance(w http.ResponseWriter, r *http.Request) {
	var postMessageSubstanceDTO dto.PostMessageSubstanceDTO
	err := json.NewDecoder(r.Body).Decode(&postMessageSubstanceDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postMessageSubstance := model.PostMessageSubstance{
		MessageSubstance: model.MessageSubstance{
			ID:   uuid.UUID{},
			Text: postMessageSubstanceDTO.Text,
		},
		PostId: postMessageSubstanceDTO.PostId,
	}

	err = handler.Service.CreatePostMessageSubstance(&postMessageSubstance)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}