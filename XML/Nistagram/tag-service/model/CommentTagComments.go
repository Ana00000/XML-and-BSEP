package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentTagComments struct {
	ID uuid.UUID `json:"id"`
	TagId uuid.UUID `json:"tag_id" gorm:"not null"`
	CommentId uuid.UUID `json:"comment_id" gorm:"not null"`
}

func(commentTagComments * CommentTagComments) BeforeCreate(scope *gorm.DB) error {
	commentTagComments.ID = uuid.New()
	return nil
}
