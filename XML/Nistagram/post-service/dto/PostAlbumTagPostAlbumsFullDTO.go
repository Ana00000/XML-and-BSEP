package dto

import "github.com/google/uuid"

type PostAlbumTagPostAlbumsFullDTO struct {
	ID uuid.UUID `json:"id"`
	TagId uuid.UUID `json:"tag_id"`
	PostAlbumId uuid.UUID `json:"post_album_id"`
}
