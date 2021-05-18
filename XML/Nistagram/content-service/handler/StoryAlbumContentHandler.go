package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type StoryAlbumContentHandler struct {
	Service * service.StoryAlbumContentService
	ContentService * service.ContentService
}

func (handler *StoryAlbumContentHandler) CreateStoryAlbumContent(w http.ResponseWriter, r *http.Request) {
	var storyAlbumContentDTO dto.StoryAlbumContentDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch storyAlbumContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	storyAlbumContent := model.StoryAlbumContent{
		Content: model.Content{
			ID:   id,
			Path: storyAlbumContentDTO.Path,
			Type: contentType,
		},
		StoryAlbumId: storyAlbumContentDTO.StoryAlbumId,
	}

	err = handler.Service.CreateStoryAlbumContent(&storyAlbumContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.ContentService.CreateContent(&storyAlbumContent.Content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
