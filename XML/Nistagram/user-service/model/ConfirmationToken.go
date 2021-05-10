package model

import (
	"github.com/google/uuid"
	"time"
)

type ConfirmationToken struct {
	ID uuid.UUID `json:"id"`
	ConfirmationToken uuid.UUID `json:"confirmation_token" gorm:"not null"`
	UserId uuid.UUID `json:"user_id" gorm:"not null"`
	CreatedTime time.Time `json:"created_time" gorm:"not null"`
	ExpiredTime time.Time `json:"expired_time" gorm:"not null"`
	IsValid bool `json:"is_valid" gorm:"not null"`
}
