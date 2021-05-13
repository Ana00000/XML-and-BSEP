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

type SinglePostContentHandler struct {
	Service * service.SinglePostContentService
	ContentService * service.ContentService
}

func (handler *SinglePostContentHandler) CreateSinglePostContent(w http.ResponseWriter, r *http.Request) {
	var singlePostContentDTO dto.SinglePostContentDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch singlePostContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	singlePostContent := model.SinglePostContent{
		Content: model.Content{
			ID:   id,
			Path: singlePostContentDTO.Path,
			Type: contentType,
		},
		SinglePostId: singlePostContentDTO.SinglePostId,
	}

	err = handler.Service.CreateSinglePostContent(&singlePostContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.ContentService.CreateContent(&singlePostContent.Content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}