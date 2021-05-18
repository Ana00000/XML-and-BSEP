package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type StoryTagHandler struct {
	Service * service.StoryTagService
	TagService * service.TagService
}

func (handler *StoryTagHandler) CreateStoryTag(w http.ResponseWriter, r *http.Request) {
	var storyTagDTO dto.StoryTagDTO
	err := json.NewDecoder(r.Body).Decode(&storyTagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	storyTag := model.StoryTag{
		Tag: model.Tag{
			ID:   id,
			Name: storyTagDTO.Name,
		},
	}

	err = handler.Service.CreateStoryTag(&storyTag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.TagService.CreateTag(&storyTag.Tag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	storyTagIDJson, _ := json.Marshal(storyTag.ID)
	w.Write(storyTagIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}