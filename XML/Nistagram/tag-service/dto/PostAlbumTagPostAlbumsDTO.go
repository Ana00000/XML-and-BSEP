package dto

import "github.com/google/uuid"

type PostAlbumTagPostAlbumsDTO struct {
	TagId uuid.UUID `json:"tag_id"`
	PostAlbumId uuid.UUID `json:"postAlbumId"`
}
