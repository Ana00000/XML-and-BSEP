package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
)

type PostAlbumContentHandler struct {
	Service * service.PostAlbumContentService
}

func (handler *PostAlbumContentHandler) CreatePostAlbumContent(w http.ResponseWriter, r *http.Request) {
	var postAlbumContentDTO dto.PostAlbumContentDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postAlbumContent := model.PostAlbumContent{
		Content: model.Content{
			ID:   uuid.UUID{},
			Path: postAlbumContentDTO.Path,
			Type: postAlbumContentDTO.Type,
		},
		PostAlbumId: postAlbumContentDTO.PostAlbumId,
	}

	err = handler.Service.CreatePostAlbumContent(&postAlbumContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
