package dto

import "github.com/google/uuid"

type ProfileSettingsPostNotificationsProfilesDTO struct{
	ProfileSettingsId uuid.UUID `json:"profile_settings_id" validate:"required"`
	PostNotificationsProfileId uuid.UUID `json:"post_notifications_profile_id" validate:"required"`
}
