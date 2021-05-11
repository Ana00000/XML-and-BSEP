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
}

func (handler *SingleStoryHandler) CreateSingleStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate,_ :=time.Parse(layout,singleStoryDTO.CreationDate)
	singleStory := model.SingleStory{
		Story : model.Story{
		ID:          uuid.UUID{},
		CreationDate: creationDate,
		UserId:      singleStoryDTO.UserId,
		LocationId:      singleStoryDTO.LocationId,
		IsDeleted:      singleStoryDTO.IsDeleted,
		Type:      singleStoryDTO.Type,
		},
	}

	err = handler.Service.CreateSingleStory(&singleStory)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

