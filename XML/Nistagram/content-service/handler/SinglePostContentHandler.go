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
}

func (handler *SinglePostContentHandler) CreateSinglePostContent(w http.ResponseWriter, r *http.Request) {
	var singlePostContentDTO dto.SinglePostContentDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	singlePostContent := model.SinglePostContent{
		Content: model.Content{
			ID:   uuid.UUID{},
			Path: singlePostContentDTO.Path,
			Type: singlePostContentDTO.Type,
		},
		SinglePostId: singlePostContentDTO.SinglePostId,
	}

	err = handler.Service.CreateSinglePostContent(&singlePostContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}