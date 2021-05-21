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

type PostCollectionPostsHandler struct {
	Service * service.PostCollectionPostsService
}

func (handler *PostCollectionPostsHandler) CreatePostCollectionPosts(w http.ResponseWriter, r *http.Request) {
	var postCollectionPostsDTO dto.PostCollectionPostsDTO
	err := json.NewDecoder(r.Body).Decode(&postCollectionPostsDTO)

	if err != nil {
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
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

