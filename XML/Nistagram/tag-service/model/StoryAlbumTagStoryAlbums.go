package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoryAlbumTagStoryAlbums struct {
	ID uuid.UUID `json:"id"`
	StoryAlbumTagId uuid.UUID `json:"storyAlbumTagId" gorm:"not null"`
	StoryAlbumId uuid.UUID `json:"storyAlbumId" gorm:"not null"`
}

func(storyAlbumTagStoryAlbums * StoryAlbumTagStoryAlbums) BeforeCreate(scope *gorm.DB) error {
	storyAlbumTagStoryAlbums.ID = uuid.New()
	return nil
}
