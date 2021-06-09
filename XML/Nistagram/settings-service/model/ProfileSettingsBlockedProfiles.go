package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettingsBlockedProfiles struct {
	ID                uuid.UUID `json:"id"`
	ProfileSettingsId uuid.UUID `json:"profile_settings_id" gorm:"not null"`
	BlockedProfileId  uuid.UUID `json:"blocked_profile_id" gorm:"not null"`
}

func (profileSettingsBlockedProfiles *ProfileSettingsBlockedProfiles) BeforeCreate(scope *gorm.DB) error {
	profileSettingsBlockedProfiles.ID = uuid.New()
	return nil
}
