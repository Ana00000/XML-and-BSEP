package dto

import (
	"github.com/google/uuid"
)

type PostAlbumFullDTO struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description"`
	CreationDate string `json:"creationDate"`
	UserID uuid.UUID `json:"userID"`
	LocationId uuid.UUID `json:"locationID"`
	IsDeleted bool `json:"isDeleted"`
}