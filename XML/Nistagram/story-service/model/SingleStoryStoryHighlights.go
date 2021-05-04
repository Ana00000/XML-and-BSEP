package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SingleStoryStoryHighlights struct {
	ID uuid.UUID `json:"id"`
	SingleStoryId uuid.UUID `json:"single_story_id" gorm:"not null"`
	StoryHighlightId uuid.UUID `json:"story_highlight_id" gorm:"not null"`
}

func(singleStoryStoryHighlights * SingleStoryStoryHighlights) BeforeCreate(scope *gorm.DB) error {
	singleStoryStoryHighlights.ID = uuid.New()
	return nil
}
