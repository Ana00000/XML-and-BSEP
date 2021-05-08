package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Advertisement struct {
	ID uuid.UUID `json:"id"`
	CampaignId uuid.UUID `json:"campaign_id" gorm:"not null"`
}

func(advertisement * Advertisement) BeforeCreate(scope *gorm.DB) error {
	advertisement.ID = uuid.New()
	return nil
}
