package dto

import "github.com/google/uuid"

type SinglePostFullDTO struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description"`
	CreationDate string `json:"creation_date"`
	UserID uuid.UUID `json:"user_id"`
	LocationId uuid.UUID `json:"location_id"`
	IsDeleted bool `json:"is_deleted"`
}
