package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	ID uuid.UUID `json:"id"`
	PostID uuid.UUID `json:"postID" gorm:"not null"`
	UserID uuid.UUID `json:"userID" gorm:"not null"`
	Liked bool `json:"liked" gorm:"not null"`
	IsFavorite bool `json:"isFavorite" gorm:"not null"`
}

func (activity *Activity) BeforeCreate(store *gorm.DB) error {
	if activity.ID == uuid.Nil {
		activity.ID = uuid.New()
	}
	return nil
}