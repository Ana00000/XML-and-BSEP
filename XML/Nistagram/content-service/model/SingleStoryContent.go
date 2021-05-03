package model

import "github.com/google/uuid"

type SingleStoryContent struct {
	Content
	SingleStoryId uuid.UUID `json:"single_story_id" gorm:"not null"`
}