package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type StoryICRHandler struct {
	Service * service.StoryICRService
}

func (handler *StoryICRHandler) CreateStoryICR(w http.ResponseWriter, r *http.Request) {
	var storyICRDTO dto.StoryICRDTO
	err := json.NewDecoder(r.Body).Decode(&storyICRDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyICR := model.StoryICR{
		InappropriateContentRequest : model.InappropriateContentRequest{
			ID:          uuid.UUID{},
			Note: 		 storyICRDTO.Note,
			UserId:      storyICRDTO.UserId,
		},
		StoryId:      storyICRDTO.StoryId,
	}

	err = handler.Service.CreateStoryICR(&storyICR)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

