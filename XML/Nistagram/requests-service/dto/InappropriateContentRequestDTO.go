package dto

import "github.com/google/uuid"

type InappropriateContentRequestDTO struct {
	Note string `json:"note"`
	UserId uuid.UUID `json:"userId"`
}