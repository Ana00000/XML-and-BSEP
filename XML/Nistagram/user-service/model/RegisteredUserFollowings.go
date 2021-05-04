package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegisteredUserFollowings struct {
	ID uuid.UUID `json:"id"`
	RegisteredUserId uuid.UUID `json:"registered_user_id" gorm:"not null"`
	FollowingUserId uuid.UUID `json:"following_user_id" gorm:"not null"`
}

func(registeredUserFollowings * RegisteredUserFollowings) BeforeCreate(scope *gorm.DB) error {
	registeredUserFollowings.ID = uuid.New()
	return nil
}
