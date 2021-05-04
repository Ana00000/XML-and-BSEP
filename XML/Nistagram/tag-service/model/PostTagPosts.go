package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostTagPosts struct {
	ID uuid.UUID `json:"id"`
	PostTagId uuid.UUID `json:"story_tag_id" gorm:"not null"`
	PostId uuid.UUID `json:"story_id" gorm:"not null"`
}

func(postTagPosts * PostTagPosts) BeforeCreate(scope *gorm.DB) error {
	postTagPosts.ID = uuid.New()
	return nil
}
