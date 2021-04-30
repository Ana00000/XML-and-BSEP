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

type SinglePostHandler struct {
	Service * service.SinglePostService
}

func (handler *SinglePostHandler) CreateSinglePost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate, _ := time.Parse(layout, singlePostDTO.CreationDate)

	singlePost := model.SinglePost{
		Post : model.Post{
			ID: uuid.UUID{},
			Description: singlePostDTO.Description,
			CreationDate: creationDate,
			UserID: singlePostDTO.UserID,
			LocationID: singlePostDTO.LocationID,
			IsDeleted: singlePostDTO.IsDeleted,
		},
		// Content
	}
	err = handler.Service.CreateSinglePost(&singlePost)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}