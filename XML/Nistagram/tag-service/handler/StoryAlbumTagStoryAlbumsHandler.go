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

type StoryAlbumTagStoryAlbumsHandler struct {
	Service *service.StoryAlbumTagStoryAlbumsService
}

func (handler *StoryAlbumTagStoryAlbumsHandler) CreateStoryAlbumTagStoryAlbums(w http.ResponseWriter, r *http.Request) {
	var storyAlbumTagStoryAlbumsDTO dto.StoryAlbumTagStoryAlbumsDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumTagStoryAlbumsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	storyAlbumTagStoryAlbums := model.StoryAlbumTagStoryAlbums{
		ID:           id,
		TagId:        storyAlbumTagStoryAlbumsDTO.TagId,
		StoryAlbumId: storyAlbumTagStoryAlbumsDTO.StoryAlbumId,
	}

	err = handler.Service.CreateStoryAlbumTagStoryAlbums(&storyAlbumTagStoryAlbums)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryAlbumTagStoryAlbumsHandler) FindAllTagsForStoryAlbumTagStoryAlbums(w http.ResponseWriter, r *http.Request) {
	var storyAlbumFullDTO []dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumFullDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForStoryAlbumTagStoryAlbums(storyAlbumFullDTO)

	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryAlbumTagStoryAlbumsHandler) FindAllTagsForStoryAlbum(w http.ResponseWriter, r *http.Request) {
	var storyAlbumFullDTO dto.StoryAlbumFullDTO
	err := json.NewDecoder(r.Body).Decode(&storyAlbumFullDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForStoryAlbum(&storyAlbumFullDTO)

	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
