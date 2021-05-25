package dto

import "github.com/google/uuid"

type ProfileSettingsMutedProfilesDTO struct{
	ProfileSettingsId uuid.UUID `json:"profile_settings_id"`
	MutedProfileId uuid.UUID `json:"muted_profile_id"`
}
