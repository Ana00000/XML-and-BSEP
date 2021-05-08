package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CampaignChosenGroup struct {
	ID uuid.UUID `json:"id"`
	CampaignId uuid.UUID `json:"campaign_id" gorm:"not null"`
	UserCategoryValue UserCategory `json:"user_category_value" gorm:"not null"`
}

func(campaignChosenGroup * CampaignChosenGroup) BeforeCreate(scope *gorm.DB) error {
	campaignChosenGroup.ID = uuid.New()
	return nil
}
