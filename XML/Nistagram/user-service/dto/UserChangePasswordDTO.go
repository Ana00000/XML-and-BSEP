package dto

import "github.com/google/uuid"

type UserChangePasswordDTO struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10,max=30"`
	ConfirmedPassword string `json:"confirmed_password" validate:"required,min=10,max=30"`
	RecoveryPasswordTokenID uuid.UUID `json:"recovery_password_token_id" validate:"required"`
}
