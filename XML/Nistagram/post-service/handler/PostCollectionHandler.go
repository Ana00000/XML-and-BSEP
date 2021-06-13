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
	"os"
	"time"
)

type PostCollectionHandler struct {
	Service * service.PostCollectionService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *PostCollectionHandler) CreatePostCollection(w http.ResponseWriter, r *http.Request) {

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionHandler",
			"action":   "CRPCL590",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-create-post-collection-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionHandler",
			"action":   "CRPCL590",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionHandler",
			"action":   "CRPCL590",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	var postCollectionDTO dto.PostCollectionDTO
	err := json.NewDecoder(r.Body).Decode(&postCollectionDTO)

	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionHandler",
			"action":   "CRPCL590",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to PostCollectionDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	postCollection := model.PostCollection{
		ID:     id,
		Title:  postCollectionDTO.Title,
		UserID: postCollectionDTO.UserID,
	}

	err = handler.Service.CreatePostCollection(&postCollection)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionHandler",
			"action":   "CRPCL590",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating post collection!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	postCollectionIDJson, _ := json.Marshal(postCollection.ID)
	w.Write(postCollectionIDJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "PostCollectionHandler",
		"action":   "CRPCL590",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created post collection!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostCollectionHandler) FindAllPostCollectionsForUserRegisteredUser(w http.ResponseWriter, r *http.Request) {

	reqUrlAuth := fmt.Sprintf("http://%s:%s/check_if_authentificated/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	response:=Request(reqUrlAuth,ExtractToken(r))
	if response.StatusCode==401{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionHandler",
			"action":   "FAPCU591",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	reqUrlAutorization := fmt.Sprintf("http://%s:%s/auth/check-find-all-post-collections-for-user-registered-user-permission/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	res := Request(reqUrlAutorization,ExtractToken(r))
	if res.StatusCode==403{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionHandler",
			"action":   "FAPCU591",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden) // 401
		return
	}

	/*if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostCollectionHandler",
			"action":   "FAPCU591",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}*/

	id := r.URL.Query().Get("id")

	postCollections := handler.Service.FindAllPostCollectionsForUserRegisteredUser(uuid.MustParse(id))
	postCollectionsJson, _ := json.Marshal(postCollections)
	if postCollectionsJson != nil {
		handler.LogInfo.WithFields(logrus.Fields{
			"status": "success",
			"location":   "PostCollectionHandler",
			"action":   "FAPCU591",
			"timestamp":   time.Now().String(),
		}).Info("Successfully found all post collections for user!")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(postCollectionsJson)
		return
	}
	handler.LogError.WithFields(logrus.Fields{
		"status": "failure",
		"location":   "PostCollectionHandler",
		"action":   "FAPCU591",
		"timestamp":   time.Now().String(),
	}).Error("Post collections for user not found!")
	w.WriteHeader(http.StatusBadRequest)
}
