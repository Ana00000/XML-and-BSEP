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

type PostAlbumHandler struct {
	Service * service.PostAlbumService
	PostService * service.PostService
}

func (handler *PostAlbumHandler) CreatePostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumDTO dto.PostAlbumDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	postAlbum := model.PostAlbum{
		Post : model.Post{
			ID: id,
			Description: postAlbumDTO.Description,
			CreationDate: time.Now(),
			UserID: postAlbumDTO.UserID,
			LocationId: postAlbumDTO.LocationID,
			IsDeleted: false,
		},
	}

	err = handler.Service.CreatePostAlbum(&postAlbum)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.PostService.CreatePost(&postAlbum.Post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	postAlbumIDJson, _ := json.Marshal(postAlbum.ID)
	w.Write(postAlbumIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}