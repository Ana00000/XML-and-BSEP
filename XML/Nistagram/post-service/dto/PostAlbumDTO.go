package dto

import (
	"github.com/google/uuid"
)

type PostAlbumDTO struct {
	Description string `json:"description"`
	CreationDate string `json:"creationDate"`
	UserID uuid.UUID `json:"userID"`
	LocationID uuid.UUID `json:"locationID"`
	IsDeleted bool `json:"isDeleted"`
}
