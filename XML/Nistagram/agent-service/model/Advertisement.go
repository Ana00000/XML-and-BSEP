package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Advertisement struct {
	ID uuid.UUID `json:"id"`
	AdvertisementContentId string `json:"advertisement_content_id"`
	CampaignRefer uuid.UUID `json:"campaign_refer"`
}

func(advertisement * Advertisement) BeforeCreate(scope *gorm.DB) error {
	advertisement.ID = uuid.New()
	return nil
}
