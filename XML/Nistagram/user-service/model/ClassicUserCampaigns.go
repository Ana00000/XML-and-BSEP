package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassicUserCampaigns struct {
	ID uuid.UUID `json:"id"`
	ClassicUserId uuid.UUID `json:"classic_user_id" gorm:"not null"`
	CampaignId uuid.UUID `json:"campaign_id" gorm:"not null"`
}

func(classicUserCampaigns * ClassicUserCampaigns) BeforeCreate(scope *gorm.DB) error {
	classicUserCampaigns.ID = uuid.New()
	return nil
}
