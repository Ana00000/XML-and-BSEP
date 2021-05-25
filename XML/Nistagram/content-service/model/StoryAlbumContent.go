package model

import "github.com/google/uuid"

type StoryAlbumContent struct {
	Content
	StoryAlbumId uuid.UUID `json:"story_album_id" gorm:"not null"`
}