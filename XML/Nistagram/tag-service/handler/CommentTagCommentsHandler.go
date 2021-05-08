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

type CommentTagCommentsHandler struct {
	Service * service.CommentTagCommentsService
}

func (handler *CommentTagCommentsHandler) CreateCommentTagComments(w http.ResponseWriter, r *http.Request) {
	var commentTagCommentsDTO dto.CommentTagCommentsDTO
	err := json.NewDecoder(r.Body).Decode(&commentTagCommentsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentTagComments := model.CommentTagComments{
		ID:          uuid.UUID{},
		CommentTagId: 		commentTagCommentsDTO.CommentTagId,
		CommentId: commentTagCommentsDTO.CommentId,
	}

	err = handler.Service.CreateCommentTagComments(&commentTagComments)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}