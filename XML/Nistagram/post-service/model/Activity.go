package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	ID uuid.UUID `json:"id"`
	PostID uuid.UUID `json:"post_id" gorm:"not null"`
	UserID uuid.UUID `json:"user_id" gorm:"not null"`
	LikedStatus LikedStatus `json:"liked_status" gorm:"not null"`
	IsFavorite bool `json:"is_favorite" gorm:"not null"`
}

func (activity *Activity) BeforeCreate(store *gorm.DB) error {
	if activity.ID == uuid.Nil {
		activity.ID = uuid.New()
	}
	return nil
}