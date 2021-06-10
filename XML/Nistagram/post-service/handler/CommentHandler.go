package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
	"time"
)

type CommentHandler struct {
	Service   *service.CommentService
	Validator *validator.Validate
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var commentDTO dto.CommentDTO
	if err := json.NewDecoder(r.Body).Decode(&commentDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to CommentDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&commentDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("CommentDTO fields aren't in the valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	layout := "2021-05-21T15:59:12.000Z"
	creationDate, _ := time.Parse(layout, commentDTO.CreationDate)

	comment := model.Comment{
		ID:           uuid.UUID{},
		CreationDate: creationDate,
		UserID:       commentDTO.UserID,
		PostID:       commentDTO.PostID,
		Text:         commentDTO.Text,
	}

	if err := handler.Service.CreateComment(&comment); err != nil {
		fmt.Println(err)
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "CommentHandler",
			"action":   "CRCOM571",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating comment!")
		w.WriteHeader(http.StatusExpectationFailed) // 417
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "CommentHandler",
		"action":   "CRCOM571",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created comment!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *CommentHandler) FindAllCommentsForPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	comments := handler.Service.FindAllCommentsForPost(uuid.MustParse(id))
	commentsJson, _ := json.Marshal(comments)
	if commentsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "CommentHandler",
			"action":   "FACFP572",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all comments for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(commentsJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "CommentHandler",
		"action":   "FACFP572",
		"timestamp":   time.Now().String(),
	}).Error("Comments for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}

func (handler *CommentHandler) FindAllUserComments(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	comments := handler.Service.FindAllUserComments(uuid.MustParse(id))
	commentsJson, _ := json.Marshal(comments)
	if commentsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "CommentHandler",
			"action":   "FAUCM573",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all user comments!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(commentsJson)
	}

	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "CommentHandler",
		"action":   "FAUCM573",
		"timestamp":   time.Now().String(),
	}).Error("All user comments not found!")
	w.WriteHeader(http.StatusBadRequest)
}
