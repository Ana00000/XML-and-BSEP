package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettingsPostNotificationsProfiles struct {
	ID                uuid.UUID `json:"id"`
	ProfileSettingsId uuid.UUID `json:"profile_settings_id" gorm:"not null"`
	PostNotificationsProfileId  uuid.UUID `json:"post_notifications_profile_id" gorm:"not null"`
}

func (profileSettingsPostNotificationsProfiles *ProfileSettingsPostNotificationsProfiles) BeforeCreate(scope *gorm.DB) error {
	profileSettingsPostNotificationsProfiles.ID = uuid.New()
	return nil
}
