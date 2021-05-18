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
	ContentService * service.ContentService
}

func (handler *PostAlbumContentHandler) CreatePostAlbumContent(w http.ResponseWriter, r *http.Request) {
	var postAlbumContentDTO dto.PostAlbumContentDTO
	err := json.NewDecoder(r.Body).Decode(&postAlbumContentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentType := model.PICTURE
	switch postAlbumContentDTO.Type {
	case "VIDEO":
		contentType = model.VIDEO
	}

	id := uuid.New()
	postAlbumContent := model.PostAlbumContent{
		Content: model.Content{
			ID:   id,
			Path: postAlbumContentDTO.Path,
			Type: contentType,
		},
		PostAlbumId: postAlbumContentDTO.PostAlbumId,
	}

	err = handler.Service.CreatePostAlbumContent(&postAlbumContent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	err = handler.ContentService.CreateContent(&postAlbumContent.Content)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
