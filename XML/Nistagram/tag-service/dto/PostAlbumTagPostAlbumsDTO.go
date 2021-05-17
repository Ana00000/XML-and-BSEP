package dto

import "github.com/google/uuid"

type PostAlbumTagPostAlbumsDTO struct {
	PostAlbumTagId uuid.UUID `json:"postAlbumTagId"`
	PostAlbumId uuid.UUID `json:"postAlbumId"`
}
