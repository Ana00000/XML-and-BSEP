package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostAlbumTagPostAlbums struct {
	ID uuid.UUID `json:"id"`
	PostAlbumTagId uuid.UUID `json:"postAlbumTagId" gorm:"not null"`
	PostAlbumId uuid.UUID `json:"postAlbumId" gorm:"not null"`
}

func(postAlbumTagPostAlbums * PostAlbumTagPostAlbums) BeforeCreate(scope *gorm.DB) error {
	postAlbumTagPostAlbums.ID = uuid.New()
	return nil
}
