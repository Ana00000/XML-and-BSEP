package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegisteredUserFollowers struct {
	ID uuid.UUID `json:"id"`
	RegisteredUserId uuid.UUID `json:"registered_user_id" gorm:"not null"`
	FollowerUserId uuid.UUID `json:"follower_user_id" gorm:"not null"`
}

func(registeredUserFollowers * RegisteredUserFollowers) BeforeCreate(scope *gorm.DB) error {
	registeredUserFollowers.ID = uuid.New()
	return nil
}