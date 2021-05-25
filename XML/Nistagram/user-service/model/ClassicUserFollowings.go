package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassicUserFollowings struct {
	ID uuid.UUID `json:"id"`
	ClassicUserId uuid.UUID `json:"classic_user_id" gorm:"not null"`
	FollowingUserId uuid.UUID `json:"following_user_id" gorm:"not null"`
}

func(classicUserFollowings * ClassicUserFollowings) BeforeCreate(scope *gorm.DB) error {
	classicUserFollowings.ID = uuid.New()
	return nil
}
