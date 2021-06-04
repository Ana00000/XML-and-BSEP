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

type StoryTagStoriesHandler struct {
	Service *service.StoryTagStoriesService
}

func (handler *StoryTagStoriesHandler) CreateStoryTagStories(w http.ResponseWriter, r *http.Request) {
	var storyTagStoriesDTO dto.StoryTagStoriesDTO
	err := json.NewDecoder(r.Body).Decode(&storyTagStoriesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storyTagStories := model.StoryTagStories{
		ID:      uuid.UUID{},
		TagId:   storyTagStoriesDTO.TagId,
		StoryId: storyTagStoriesDTO.StoryId,
	}

	err = handler.Service.CreateStoryTagStories(&storyTagStories)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryTagStoriesHandler) FindAllTagsForStory(w http.ResponseWriter, r *http.Request) {
	var singleStoryDTO dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoryDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForStory(&singleStoryDTO)

	contentsForStoriesJson, _ := json.Marshal(tags)
	w.Write(contentsForStoriesJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryTagStoriesHandler) FindAllTagsForStories(w http.ResponseWriter, r *http.Request) {
	var singleStoriesDTO []dto.SingleStoryDTO
	err := json.NewDecoder(r.Body).Decode(&singleStoriesDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForStories(singleStoriesDTO)

	contentsForStoriesJson, _ := json.Marshal(tags)
	w.Write(contentsForStoriesJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *StoryTagStoriesHandler) FindStoryTagStoriesForStoryId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	storyTagStories := handler.Service.FindStoryTagStoriesForStoryId(uuid.MustParse(id))
	storyTagStoriesJson, _ := json.Marshal(storyTagStories)
	if storyTagStoriesJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(storyTagStoriesJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}
