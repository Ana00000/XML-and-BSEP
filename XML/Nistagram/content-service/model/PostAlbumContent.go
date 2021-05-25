package model

import "github.com/google/uuid"

type PostAlbumContent struct {
	Content
	PostAlbumId uuid.UUID `json:"post_album_id" gorm:"not null"`
}