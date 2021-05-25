package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettingsRejectedMessageProfiles struct {
	ID uuid.UUID `json:"id"`
	ProfileSettingsId uuid.UUID `json:"profile_settings_id" gorm:"not null"`
	RejectedMessageProfileId uuid.UUID `json:"rejected_message_profile_id" gorm:"not null"`
}

func(profileSettingsRejectedMessageProfiles * ProfileSettingsRejectedMessageProfiles) BeforeCreate(scope *gorm.DB) error {
	profileSettingsRejectedMessageProfiles.ID = uuid.New()
	return nil
}
