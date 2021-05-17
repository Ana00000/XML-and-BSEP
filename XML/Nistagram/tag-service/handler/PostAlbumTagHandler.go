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

type PostAlbumTagHandler struct {
	Service * service.PostAlbumTagService
	TagService * service.TagService
}

func (handler *PostAlbumTagHandler) CreatePostAlbumTag(w http.ResponseWriter, r *http.Request) {
	var postAlbumTagDTO dto.PostAlbumTagDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumTagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	postAlbumTag := model.PostAlbumTag{
		Tag: model.Tag{
			ID:   id,
			Name: postAlbumTagDTO.Name,
		},
	}

	err = handler.Service.CreatePostAlbumTag(&postAlbumTag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.TagService.CreateTag(&postAlbumTag.Tag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	postAlbumTagIDJson, _ := json.Marshal(postAlbumTag.ID)
	w.Write(postAlbumTagIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
