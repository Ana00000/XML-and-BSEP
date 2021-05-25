package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostAlbumTagPostAlbums struct {
	ID uuid.UUID `json:"id"`
	TagId uuid.UUID `json:"tag_id" gorm:"not null"`
	PostAlbumId uuid.UUID `json:"post_album_id" gorm:"not null"`
}

func(postAlbumTagPostAlbums * PostAlbumTagPostAlbums) BeforeCreate(scope *gorm.DB) error {
	postAlbumTagPostAlbums.ID = uuid.New()
	return nil
}
