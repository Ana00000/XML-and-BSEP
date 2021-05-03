package model

import "github.com/google/uuid"

type AdvertisementContent struct {
	Content
	Link            string    `json:"link" gorm:"not null"`
	AdvertisementId uuid.UUID `json:"advertisement_id" gorm:"not null"`
}