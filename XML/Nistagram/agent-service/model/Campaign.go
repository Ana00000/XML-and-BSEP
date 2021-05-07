package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	ID uuid.UUID `json:"id"`
	ExposureTime time.Time `json:"exposure_time" gorm:"not null"`
}

func(campaign * Campaign) BeforeCreate(scope *gorm.DB) error {
	campaign.ID = uuid.New()
	return nil
}