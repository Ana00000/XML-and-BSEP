package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"net/http"
)

type PostCollectionPostsHandler struct {
	Service *service.PostCollectionPostsService
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

func (handler *PostCollectionPostsHandler) FindAllPostCollectionPostsForPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	postCollectionPosts := handler.Service.FindAllPostCollectionPostsForPost(uuid.MustParse(id))
	postCollectionPostsJson, _ := json.Marshal(postCollectionPosts)
	if postCollectionPostsJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(postCollectionPostsJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}
