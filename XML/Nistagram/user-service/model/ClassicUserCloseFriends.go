package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassicUserCloseFriends struct {
	ID uuid.UUID `json:"id"`
	ClassicUserId uuid.UUID `json:"classic_user_id" gorm:"not null"`
	CloseFriendUserId uuid.UUID `json:"close_friend_user_id" gorm:"not null"`
}

func(classicUserCloseFriends * ClassicUserCloseFriends) BeforeCreate(scope *gorm.DB) error {
	classicUserCloseFriends.ID = uuid.New()
	return nil
}