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

	layout := "2006-01-02T15:04:05.000Z"
	creationDate,_ :=time.Parse(layout,storyAlbumDTO.CreationDate)
	storyAlbum := model.StoryAlbum{
		Story : model.Story{
			ID:          uuid.UUID{},
			CreationDate: creationDate,
			UserId:      storyAlbumDTO.UserId,
			LocationId:      storyAlbumDTO.LocationId,
			IsDeleted:      storyAlbumDTO.IsDeleted,
			Type:      storyAlbumDTO.Type,
			//StoryICRs: nil,
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

