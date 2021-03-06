package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoryHighlight struct {
	ID uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId" gorm:"not null"`
	Title string `json:"title" gorm:"not null"`
}

func(storyHighlight * StoryHighlight) BeforeCreate(scope *gorm.DB) error {
	storyHighlight.ID = uuid.New()
	return nil
}