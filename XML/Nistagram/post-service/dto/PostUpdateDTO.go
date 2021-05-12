package dto

import "github.com/google/uuid"

type PostUpdateDTO struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description"`
	LocationID uuid.UUID `json:"locationID"`
}

