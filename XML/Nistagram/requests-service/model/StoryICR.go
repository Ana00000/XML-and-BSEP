package model

import "github.com/google/uuid"

type StoryICR struct {
	InappropriateContentRequest
	StoryId uuid.UUID `json:"storyId" gorm:"not null"`
}
