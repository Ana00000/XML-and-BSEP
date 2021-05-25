package dto

import "github.com/google/uuid"

type UserTagFullDTO struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	UserId uuid.UUID `json:"user_id"`
}
