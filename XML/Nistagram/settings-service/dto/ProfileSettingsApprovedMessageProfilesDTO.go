package dto

import "github.com/google/uuid"

type ProfileSettingsApprovedMessageProfilesDTO struct {
	ProfileSettingsId        uuid.UUID `json:"profile_settings_id" validate:"required"`
	ApprovedMessageProfileId uuid.UUID `json:"approved_message_profile_id" validate:"required"`
}
