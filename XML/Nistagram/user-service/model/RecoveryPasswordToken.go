package model

import (
	"github.com/google/uuid"
	"time"
)

type RecoveryPasswordToken struct {
	ID uuid.UUID `json:"id"`
	RecoveryPasswordToken uuid.UUID `json:"recovery_password_token" gorm:"not null"`
	UserId uuid.UUID `json:"user_id" gorm:"not null"`
	CreatedTime time.Time `json:"created_time" gorm:"not null"`
	ExpirationTime time.Time `json:"expiration_time" gorm:"not null"`
	IsValid bool `json:"is_valid" gorm:"not null"`
}
