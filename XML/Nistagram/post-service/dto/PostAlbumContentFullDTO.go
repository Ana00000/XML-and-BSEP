package dto

import "github.com/google/uuid"

type PostAlbumContentFullDTO struct {
	ID uuid.UUID `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
	PostAlbumId uuid.UUID `json:"post_album_id"`
}
