package handler

import (
	"../dto"
	"../model"
	"../service"
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

	postCollection := model.PostCollection{
		ID: uuid.UUID{},
		Title: postCollectionDTO.Title,
		UserID: postCollectionDTO.UserID,
	}

	err = handler.Service.CreatePostCollection(&postCollection)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
