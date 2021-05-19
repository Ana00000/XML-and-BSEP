package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
	_ "strconv"
)

type StoryAlbumTagHandler struct {
	Service * service.StoryAlbumTagService
	TagService * service.TagService
}

func (handler *StoryAlbumTagHandler) CreateStoryAlbumTag(w http.ResponseWriter, r *http.Request) {
	var storyAlbumTagDTO dto.StoryAlbumTagDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumTagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	storyAlbumTag := model.StoryAlbumTag{
		Tag: model.Tag{
			ID:   id,
			Name: storyAlbumTagDTO.Name,
		},
	}

	err = handler.Service.CreateStoryAlbumTag(&storyAlbumTag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.TagService.CreateTag(&storyAlbumTag.Tag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	storyAlbumTagIDJson, _ := json.Marshal(storyAlbumTag.ID)
	w.Write(storyAlbumTagIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
