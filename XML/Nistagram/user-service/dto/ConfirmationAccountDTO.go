package dto

import "github.com/google/uuid"

type ConfirmationAccountDTO struct {
	ConfirmationToken uuid.UUID `json:"confirmation_token"`
	UserId uuid.UUID `json:"user_id"`
}
