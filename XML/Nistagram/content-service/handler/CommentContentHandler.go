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

type CommentContentHandler struct {
	Service * service.CommentContentService
}

func (handler *CommentContentHandler) CreateCommentContent(w http.ResponseWriter, r *http.Request) {
	var commentContentDTO dto.CommentContentDTO
	err := json.NewDecoder(r.Body).Decode(&commentContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentContent := model.CommentContent{
		Content: model.Content{
			ID:   uuid.UUID{},
			Path: commentContentDTO.Path,
			Type: commentContentDTO.Type,
		},
		CommentId: commentContentDTO.CommentId,
	}

	err = handler.Service.CreateCommentContent(&commentContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
