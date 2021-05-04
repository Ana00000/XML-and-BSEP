package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type CommentHandler struct {
	Service * service.CommentService
}

func (handler *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var commentDTO dto.CommentDTO
	err := json.NewDecoder(r.Body).Decode(&commentDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate, _ := time.Parse(layout, commentDTO.CreationDate)

	comment := model.Comment{
		ID: uuid.UUID{},
		//ContentID: commentDTO.ContentID,
		CreationDate: creationDate,
		UserID: commentDTO.UserID,
		PostID: commentDTO.PostID,
		//CommentICRs: nil,
	}

	err = handler.Service.CreateComment(&comment)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}