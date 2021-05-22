package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoryAlbumTagStoryAlbums struct {
	ID uuid.UUID `json:"id"`
	TagId uuid.UUID `json:"tag_id" gorm:"not null"`
	StoryAlbumId uuid.UUID `json:"story_album_id" gorm:"not null"`
}

func(storyAlbumTagStoryAlbums * StoryAlbumTagStoryAlbums) BeforeCreate(scope *gorm.DB) error {
	storyAlbumTagStoryAlbums.ID = uuid.New()
	return nil
}
