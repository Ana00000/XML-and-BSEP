package dto

import "github.com/google/uuid"

type ProfileSettingsBlockedProfilesDTO struct{
	ProfileSettingsId uuid.UUID `json:"profile_settings_id"`
	BlockedProfileId uuid.UUID `json:"blocked_profile_id"`
}