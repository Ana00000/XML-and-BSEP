package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
	_ "strconv"
)

type PostTagPostsHandler struct {
	Service * service.PostTagPostsService
}

func (handler *PostTagPostsHandler) CreatePostTagPosts(w http.ResponseWriter, r *http.Request) {
	var postTagPostsDTO dto.PostTagPostsDTO
	err := json.NewDecoder(r.Body).Decode(&postTagPostsDTO)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postTagPosts := model.PostTagPosts{
		ID:        uuid.UUID{},
		PostTagId: postTagPostsDTO.PostTagId,
		PostId:    postTagPostsDTO.PostId,
	}

	err = handler.Service.CreatePostTagPosts(&postTagPosts)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostTagPostsHandler) FindAllTagsForPost(w http.ResponseWriter, r *http.Request) {
	var singlePostDTO dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForPost(&singlePostDTO)

	tagsForPostJson, _ := json.Marshal(tags)
	w.Write(tagsForPostJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostTagPostsHandler) FindAllTagsForPosts(w http.ResponseWriter, r *http.Request) {
	var singlePostsDTO []dto.SinglePostDTO
	err := json.NewDecoder(r.Body).Decode(&singlePostsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tags = handler.Service.FindAllTagsForPosts(singlePostsDTO)

	tagsForPostsJson, _ := json.Marshal(tags)
	w.Write(tagsForPostsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}