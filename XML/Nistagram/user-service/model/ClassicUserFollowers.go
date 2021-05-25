package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassicUserFollowers struct {
	ID uuid.UUID `json:"id"`
	ClassicUserId uuid.UUID `json:"classic_user_id" gorm:"not null"`
	FollowerUserId uuid.UUID `json:"follower_user_id" gorm:"not null"`
}

func(classicUserFollowers * ClassicUserFollowers) BeforeCreate(scope *gorm.DB) error {
	classicUserFollowers.ID = uuid.New()
	return nil
}