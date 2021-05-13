package dto

import (
	"github.com/google/uuid"
)

type PostAlbumContentDTO struct {
	Path string `json:"path"`
	Type string `json:"type"`
	PostAlbumId uuid.UUID `json:"post_album_id"`
}
