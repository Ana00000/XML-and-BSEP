package dto

import "github.com/google/uuid"

type VerifiedReturnDTO struct {
	UserEmail string `json:"user_email"`
	RecoveryPasswordTokenID uuid.UUID `json:"recovery_password_token_id"`
}
