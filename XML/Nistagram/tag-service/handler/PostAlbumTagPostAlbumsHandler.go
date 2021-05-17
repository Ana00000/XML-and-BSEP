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

type PostAlbumTagPostAlbumsHandler struct {
	Service * service.PostAlbumTagPostAlbumsService
}

func (handler *PostAlbumTagPostAlbumsHandler) CreatePostAlbumTagPostAlbums(w http.ResponseWriter, r *http.Request) {
	var postAlbumTagPostAlbumsDTO dto.PostAlbumTagPostAlbumsDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumTagPostAlbumsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	postAlbumTagPostAlbums := model.PostAlbumTagPostAlbums{
		ID:        			id,
		PostAlbumTagId: 	postAlbumTagPostAlbumsDTO.PostAlbumTagId,
		PostAlbumId:    	postAlbumTagPostAlbumsDTO.PostAlbumId,
	}

	err = handler.Service.CreatePostAlbumTagPostAlbums(&postAlbumTagPostAlbums)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}