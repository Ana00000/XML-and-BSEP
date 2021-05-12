package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/google/uuid"
)

type PostAlbumContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	PostAlbumId uuid.UUID `json:"post_album_id"`
}
