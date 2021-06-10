package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettingsMutedProfiles struct {
	ID                uuid.UUID `json:"id"`
	ProfileSettingsId uuid.UUID `json:"profile_settings_id" gorm:"not null"`
	MutedProfileId    uuid.UUID `json:"muted_profile_id" gorm:"not null"`
}

func (profileSettingsMutedProfiles *ProfileSettingsMutedProfiles) BeforeCreate(scope *gorm.DB) error {
	profileSettingsMutedProfiles.ID = uuid.New()
	return nil
}
