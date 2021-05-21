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

	var findTag = handler.TagService.FindTagByName(storyTagDTO.Name)
	var storyTag model.StoryTag

	if findTag == nil {
		id := uuid.New()
		storyTag = model.StoryTag{
			Tag: model.Tag{
				ID:   id,
				Name: storyTagDTO.Name,
			},
		}

		if err := handler.Service.CreateStoryTag(&storyTag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

		if err := handler.TagService.CreateTag(&storyTag.Tag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	} else {
		id := uuid.New()
		storyTag = model.StoryTag{
			Tag: model.Tag{
				ID:   id,
				Name: findTag.Name,
			},
		}

		if err := handler.Service.CreateStoryTag(&storyTag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}

	storyTagIDJson, _ := json.Marshal(storyTag.ID)
	w.Write(storyTagIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}