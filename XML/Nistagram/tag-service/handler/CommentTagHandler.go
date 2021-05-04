package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type CommentTagHandler struct {
	Service * service.CommentTagService
}

func (handler *CommentTagHandler) CreateCommentTag(w http.ResponseWriter, r *http.Request) {
	var commentTagDTO dto.CommentTagDTO
	err := json.NewDecoder(r.Body).Decode(&commentTagDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentTag := model.CommentTag{
		Tag: model.Tag{
			ID:          uuid.UUID{},
			Name: 		commentTagDTO.Name,
		},
	}

	err = handler.Service.CreateCommentTag(&commentTag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
