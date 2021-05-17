package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
	"time"
)

type StoryAlbumHandler struct {
	Service * service.StoryAlbumService
}

func (handler *StoryAlbumHandler) CreateStoryAlbum(w http.ResponseWriter, r *http.Request) {
	var storyAlbumDTO dto.StoryAlbumDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyAlbumType := model.CLOSE_FRIENDS
	switch storyAlbumDTO.Type {
	case "ALL_FRIENDS":
		storyAlbumType = model.ALL_FRIENDS
	case "PUBLIC":
		storyAlbumType = model.PUBLIC
	}

	id := uuid.New()
	storyAlbum := model.StoryAlbum{
		Story : model.Story{
			ID:          	id,
			CreationDate: 	time.Now(),
			Description:    storyAlbumDTO.Description,
			UserId:      	storyAlbumDTO.UserId,
			LocationId:     storyAlbumDTO.LocationId,
			IsDeleted:      false,
			Type:      		storyAlbumType,
		},
	}

	err = handler.Service.CreateStoryAlbum(&storyAlbum)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

