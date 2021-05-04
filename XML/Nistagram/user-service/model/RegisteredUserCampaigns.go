package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegisteredUserCampaigns struct {
	ID uuid.UUID `json:"id"`
	RegisteredUserId uuid.UUID `json:"registered_user_id" gorm:"not null"`
	CampaignId uuid.UUID `json:"campaign_id" gorm:"not null"`
}

func(registeredUserCampaigns * RegisteredUserCampaigns) BeforeCreate(scope *gorm.DB) error {
	registeredUserCampaigns.ID = uuid.New()
	return nil
}
