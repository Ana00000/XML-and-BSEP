package dto

import "github.com/google/uuid"

type ProfileSettingsStoryNotificationsProfilesDTO struct{
	ProfileSettingsId uuid.UUID `json:"profile_settings_id" validate:"required"`
	StoryNotificationsProfileId uuid.UUID `json:"story_notifications_profile_id" validate:"required"`
}