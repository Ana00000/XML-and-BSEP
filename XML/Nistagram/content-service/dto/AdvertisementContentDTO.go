package dto

import (
	"github.com/google/uuid"
)

type AdvertisementContentDTO struct {
	Path string `json:"path"`
	Type string `json:"type"`
	Link string `json:"link"`
	AdvertisementId uuid.UUID `json:"advertisement_id"`
}
