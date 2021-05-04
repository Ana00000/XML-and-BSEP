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

type PostAlbumHandler struct {
	Service * service.PostAlbumService
}

func (handler *PostAlbumHandler) CreatePostAlbum(w http.ResponseWriter, r *http.Request) {
	var postAlbumDTO dto.PostAlbumDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	creationDate, _ := time.Parse(layout, postAlbumDTO.CreationDate)

	postAlbum := model.PostAlbum{
		Post : model.Post{
			ID: uuid.UUID{},
			Description: postAlbumDTO.Description,
			CreationDate: creationDate,
			UserID: postAlbumDTO.UserID,
			LocationId: postAlbumDTO.LocationID,
			IsDeleted: postAlbumDTO.IsDeleted,
		},
	}
	err = handler.Service.CreatePostAlbum(&postAlbum)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}