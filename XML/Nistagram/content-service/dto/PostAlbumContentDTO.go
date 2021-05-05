package dto

import (
	"../model"
	"github.com/google/uuid"
)

type PostAlbumContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	PostAlbumId uuid.UUID `json:"post_album_id"`
}
