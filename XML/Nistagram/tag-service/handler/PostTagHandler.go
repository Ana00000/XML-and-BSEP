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

	id := uuid.New()
	postTag := model.PostTag{
		Tag: model.Tag{
			ID:   id,
			Name: postTagDTO.Name,
		},
	}

	err = handler.Service.CreatePostTag(&postTag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.TagService.CreateTag(&postTag.Tag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
