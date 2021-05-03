package model

import (
	user "../../user-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	ID uuid.UUID `json:"id"`
	Advertisements []Advertisement `json:"advertisements" gorm:"foreignKey:CampaignId"`
	ExposureTime time.Time `json:"exposure_time" gorm:"not null"`
	ChosenGroups []user.UserCategory `gorm:"many2many:campaign_chosen_groups"`
}

func(campaign * Campaign) BeforeCreate(scope *gorm.DB) error {
	campaign.ID = uuid.New()
	return nil
}