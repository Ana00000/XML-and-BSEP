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

type StoryHandler struct {
	Service * service.StoryService
}

func (handler *StoryHandler) CreateStory(w http.ResponseWriter, r *http.Request) {
	var storyDTO dto.StoryDTO
	err := json.NewDecoder(r.Body).Decode(&storyDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate,_ :=time.Parse(layout,storyDTO.CreationDate)
	story := model.Story{
		ID:          uuid.UUID{},
		CreationDate: creationDate,
		UserId:      storyDTO.UserId,
		LocationId: storyDTO.LocationId,
		IsDeleted:      storyDTO.IsDeleted,
		Type:      storyDTO.Type,
	}

	err = handler.Service.CreateStory(&story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

