package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettingsStoryNotificationsProfiles struct {
	ID                uuid.UUID `json:"id"`
	ProfileSettingsId uuid.UUID `json:"profile_settings_id" gorm:"not null"`
	StoryNotificationsProfileId  uuid.UUID `json:"story_notifications_profile_id" gorm:"not null"`
}

func (profileSettingsStoryNotificationsProfiles *ProfileSettingsStoryNotificationsProfiles) BeforeCreate(scope *gorm.DB) error {
	profileSettingsStoryNotificationsProfiles.ID = uuid.New()
	return nil
}

