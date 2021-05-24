package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoryTagStories struct {
	ID uuid.UUID `json:"id"`
	TagId uuid.UUID `json:"tag_id" gorm:"not null"`
	StoryId uuid.UUID `json:"story_id" gorm:"not null"`
}

func(storyTagStories * StoryTagStories) BeforeCreate(scope *gorm.DB) error {
	storyTagStories.ID = uuid.New()
	return nil
}
