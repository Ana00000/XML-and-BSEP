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

type SingleStoryHandler struct {
	Service * service.SingleStoryService
	StoryService * service.StoryService
}

func (handler *SingleStoryHandler) CreateSingleStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	singleStoryType := model.CLOSE_FRIENDS
	switch singleStoryDTO.Type {
	case "ALL_FRIENDS":
		singleStoryType = model.ALL_FRIENDS
	case "PUBLIC":
		singleStoryType = model.PUBLIC
	}

	id := uuid.New()
	singleStory := model.SingleStory{
		Story: model.Story{
			ID:           id,
			CreationDate: time.Now(),
			Description:  singleStoryDTO.Description,
			UserId:       singleStoryDTO.UserId,
			LocationId:   singleStoryDTO.LocationId,
			IsDeleted:    false,
			Type:         singleStoryType,
		},
	}

	err = handler.Service.CreateSingleStory(&singleStory)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.StoryService.CreateStory(&singleStory.Story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	singleStoryIDJson, _ := json.Marshal(singleStory.ID)
	w.Write(singleStoryIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}