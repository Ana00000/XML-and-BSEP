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

type PostHandler struct {
	Service *service.PostService
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var postDTO dto.PostDTO
	err := json.NewDecoder(r.Body).Decode(&postDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate, _ := time.Parse(layout, postDTO.CreationDate)

	post := model.Post{
		ID:           uuid.UUID{},
		Description:  postDTO.Description,
		CreationDate: creationDate,
		UserID:       postDTO.UserID,
		Activities: nil,
		Comments: nil,
		LocationID: postDTO.LocationID,
		IsDeleted: postDTO.IsDeleted,
		PostICRs: nil,
	}

	err = handler.Service.CreatePost(&post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
