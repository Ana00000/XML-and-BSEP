package dto

import "github.com/google/uuid"

type RecoveryPasswordDTO struct {
	RecoveryPasswordToken uuid.UUID `json:"recovery_password_token"`
	UserId uuid.UUID `json:"user_id"`
}
