package dto

import (
	"github.com/google/uuid"
	"time"
)

type SelectedStoryAlbumDTO struct {
	Path []string `json:"path"`
	Description string `json:"description"`
	CreationDate time.Time `json:"creation_date"`
	UserId uuid.UUID `json:"user_id"`
	StoryAlbumId uuid.UUID `json:"story_album_id"`
	LocationId uuid.UUID `json:"location_id"`
	Country string `json:"country"`
	City string `json:"city"`
	StreetName string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	Tags []string `json:"tags"`
	IsDeleted bool `json:"is_deleted"`
	Type []string `json:"type"`
}
