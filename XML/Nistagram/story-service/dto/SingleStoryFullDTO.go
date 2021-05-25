package dto

import "github.com/google/uuid"

type SingleStoryFullDTO struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description"`
	CreationDate string `json:"creation_date"`
	UserId uuid.UUID `json:"user_id"`
	LocationId uuid.UUID `json:"location_id"`
	IsDeleted bool `json:"is_deleted"`
	IsExpired bool `json:"is_expired"`
	Type string `json:"type"`
}
