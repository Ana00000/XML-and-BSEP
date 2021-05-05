package dto

import "github.com/google/uuid"

type ProfileSettingsApprovedMessageProfilesDTO struct{
	ProfileSettingsId uuid.UUID `json:"profile_settings_id"`
	ApprovedMessageProfileId uuid.UUID `json:"approved_message_profile_id"`
}
