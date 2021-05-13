package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type SinglePostHandler struct {
	Service * service.SinglePostService
	PostService * service.PostService
}

func (handler *SinglePostHandler) CreateSinglePost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	singlePost := model.SinglePost{
		Post : model.Post{
			ID: uuid.New(),
			Description: singlePostDTO.Description,
			CreationDate: time.Now(),
			UserID: singlePostDTO.UserID,
			LocationId: singlePostDTO.LocationID,
			IsDeleted: false,
		},
	}

	err = handler.Service.CreateSinglePost(&singlePost)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.PostService.CreatePost(&singlePost.Post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}