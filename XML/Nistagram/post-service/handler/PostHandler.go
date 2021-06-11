package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
	"time"
)

type PostHandler struct {
	PostService *service.PostService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var postDTO dto.PostDTO
	err := json.NewDecoder(r.Body).Decode(&postDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostHandler",
			"action":   "CPOST530",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post := model.Post{
		ID:           uuid.UUID{},
		Description:  postDTO.Description,
		CreationDate: time.Now(),
		UserID:       postDTO.UserID,
		LocationId:   postDTO.LocationID,
		IsDeleted:    false,
	}

	err = handler.PostService.CreatePost(&post)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostHandler",
			"action":   "CPOST530",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating post!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostHandler",
		"action":   "CPOST530",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created post!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var postDTO dto.PostUpdateDTO

	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostHandler",
			"action":   "UPOST531",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.PostService.UpdatePost(&postDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostHandler",
			"action":   "UPOST531",
			"timestamp":   time.Now().String(),
		}).Error("Failed updating post!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostHandler",
		"action":   "UPOST531",
		"timestamp":   time.Now().String(),
	}).Info("Successfully updated post!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
