package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID uuid.UUID `json:"id"`
	CreationDate time.Time `json:"creation_date" gorm:"not null"`
	UserID uuid.UUID `json:"user_id" gorm:"not null"`
	PostID uuid.UUID `json:"post_id" gorm:"not null"`
	Text string `json:"text" gorm:"not null"`
}

func (comment *Comment) BeforeCreate(scope *gorm.DB) error {
	if comment.ID == uuid.Nil {
		comment.ID = uuid.New()
	}
	return nil
}