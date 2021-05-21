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

type PostTagHandler struct {
	Service * service.PostTagService
	TagService * service.TagService
}

func (handler *PostTagHandler) CreatePostTag(w http.ResponseWriter, r *http.Request) {
	var postTagDTO dto.PostTagDTO
	err := json.NewDecoder(r.Body).Decode(&postTagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var findTag = handler.TagService.FindTagByName(postTagDTO.Name)
	var postTag model.PostTag

	if findTag == nil {
		id := uuid.New()
		postTag = model.PostTag{
			Tag: model.Tag{
				ID:   id,
				Name: postTagDTO.Name,
			},
		}

		if err := handler.Service.CreatePostTag(&postTag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

		if err := handler.TagService.CreateTag(&postTag.Tag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	} else {
		id := uuid.New()
		postTag = model.PostTag{
			Tag: model.Tag{
				ID:   id,
				Name: findTag.Name,
			},
		}

		if err := handler.Service.CreatePostTag(&postTag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}

	postTagIDJson, _ := json.Marshal(postTag.ID)
	w.Write(postTagIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
