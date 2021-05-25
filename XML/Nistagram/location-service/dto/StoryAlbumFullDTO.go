package dto

import (
	"github.com/google/uuid"
	"time"
)

type StoryAlbumFullDTO struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description"`
	CreationDate time.Time `json:"creationDate"`
	UserId uuid.UUID `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	IsExpired bool `json:"isExpired"`
	Type string `json:"type"`
}
