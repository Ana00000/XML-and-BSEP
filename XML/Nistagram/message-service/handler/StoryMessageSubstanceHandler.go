package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/service"
	"net/http"
)

type StoryMessageSubstanceHandler struct {
	Service *service.StoryMessageSubstanceService
}

func (handler *StoryMessageSubstanceHandler) CreateStoryMessageSubstance(w http.ResponseWriter, r *http.Request) {
	var storyMessageSubstanceDTO dto.StoryMessageSubstanceDTO
	err := json.NewDecoder(r.Body).Decode(&storyMessageSubstanceDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyMessageSubstance := model.StoryMessageSubstance{
		MessageSubstance: model.MessageSubstance{
			ID:   uuid.UUID{},
			Text: storyMessageSubstanceDTO.Text,
		},
		StoryId: storyMessageSubstanceDTO.StoryId,
	}

	err = handler.Service.CreateStoryMessageSubstance(&storyMessageSubstance)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
