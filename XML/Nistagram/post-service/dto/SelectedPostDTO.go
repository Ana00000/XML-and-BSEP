package dto

import (
	"github.com/google/uuid"
	"time"
)

type SelectedPostDTO struct {
	Path string `json:"path"`
	Description string `json:"description"`
	CreationDate time.Time `json:"creation_date"`
	UserId uuid.UUID `json:"user_id"`
	PostId uuid.UUID `json:"post_id"`
	LocationId uuid.UUID `json:"location_id"`
	Country string `json:"country"`
	City string `json:"city"`
	StreetName string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	Tags []string `json:"tags"`
	IsDeleted bool `json:"is_deleted"`
}
