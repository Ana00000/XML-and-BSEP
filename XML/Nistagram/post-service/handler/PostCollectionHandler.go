package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type PostCollectionHandler struct {
	Service * service.PostCollectionService
}

func (handler *PostCollectionHandler) CreatePostCollection(w http.ResponseWriter, r *http.Request) {
	var postCollectionDTO dto.PostCollectionDTO
	err := json.NewDecoder(r.Body).Decode(&postCollectionDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	postCollection := model.PostCollection{
		ID: id,
		Title: postCollectionDTO.Title,
		UserID: postCollectionDTO.UserID,
	}

	err = handler.Service.CreatePostCollection(&postCollection)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	postCollectionIDJson, _ := json.Marshal(postCollection.ID)
	w.Write(postCollectionIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
