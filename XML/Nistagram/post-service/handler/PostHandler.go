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

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var postDTO dto.PostDTO
	err := json.NewDecoder(r.Body).Decode(&postDTO)

	if err != nil {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var postDTO dto.PostUpdateDTO

	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.PostService.UpdatePost(&postDTO)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}


