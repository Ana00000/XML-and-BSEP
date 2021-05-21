package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
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
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	layout := "2021-05-21T15:59:12.000Z"
	creationDate, _ := time.Parse(layout, commentDTO.CreationDate)

	comment := model.Comment{
		ID: uuid.UUID{},
		CreationDate: creationDate,
		UserID: commentDTO.UserID,
		PostID: commentDTO.PostID,
	}

	if err = handler.Service.CreateComment(&comment); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CommentHandler) FindAllCommentsForPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	comments := handler.Service.FindAllCommentsForPost(uuid.MustParse(id))
	commentsJson, _ := json.Marshal(comments)
	if commentsJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(commentsJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *CommentHandler) FindAllUserComments(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	comments := handler.Service.FindAllUserComments(uuid.MustParse(id))
	commentsJson, _ := json.Marshal(comments)
	if commentsJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(commentsJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}
