package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
	"time"
)

type PostCollectionPostsHandler struct {
	Service * service.PostCollectionPostsService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *PostCollectionPostsHandler) CreatePostCollectionPosts(w http.ResponseWriter, r *http.Request) {
	var postCollectionPostsDTO dto.PostCollectionPostsDTO
	err := json.NewDecoder(r.Body).Decode(&postCollectionPostsDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionPostsHandler",
			"action":   "CRPCP690",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostCollectionPostsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	postCollectionPosts := model.PostCollectionPosts{
		ID:               id,
		PostCollectionId: postCollectionPostsDTO.PostCollectionId,
		SinglePostId:     postCollectionPostsDTO.SinglePostId,
	}
	err = handler.Service.CreatePostCollectionPosts(&postCollectionPosts)
	if err != nil {
		fmt.Println(err)
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionPostsHandler",
			"action":   "CRPCP690",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating post collection posts!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostCollectionPostsHandler",
		"action":   "CRPCP690",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created post collection posts!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostCollectionPostsHandler) FindAllPostCollectionPostsForPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	postCollectionPosts := handler.Service.FindAllPostCollectionPostsForPost(uuid.MustParse(id))
	postCollectionPostsJson, _ := json.Marshal(postCollectionPosts)
	if postCollectionPostsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "PostCollectionPostsHandler",
			"action":   "FAPCP691",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all post collection posts for post!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(postCollectionPostsJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "PostCollectionPostsHandler",
		"action":   "FAPCP691",
		"timestamp":   time.Now().String(),
	}).Error("Post collection posts for post not found!")
	w.WriteHeader(http.StatusBadRequest)
}
