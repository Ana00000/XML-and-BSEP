package dto

import "github.com/google/uuid"

type InappropriateContentRequestDTO struct {
	Note   string    `json:"note" validate:"required"`
	UserId uuid.UUID `json:"userId" validate:"required"`
}
