package dto

import "github.com/google/uuid"

type LogInResponseDTO struct {
	ID uuid.UUID `json:"id"`
	Token string `json:"token"`
}
