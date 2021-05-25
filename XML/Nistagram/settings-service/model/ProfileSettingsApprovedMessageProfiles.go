package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettingsApprovedMessageProfiles struct {
	ID uuid.UUID `json:"id"`
	ProfileSettingsId uuid.UUID `json:"profile_settings_id" gorm:"not null"`
	ApprovedMessageProfileId uuid.UUID `json:"approved_message_profile_id" gorm:"not null"`
}

func(profileSettingsApprovedMessageProfiles * ProfileSettingsApprovedMessageProfiles) BeforeCreate(scope *gorm.DB) error {
	profileSettingsApprovedMessageProfiles.ID = uuid.New()
	return nil
}
