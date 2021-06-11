package dto

import "github.com/google/uuid"

type ProfileSettingsRejectedMessageProfilesDTO struct {
	ProfileSettingsId        uuid.UUID `json:"profile_settings_id" validate:"required"`
	RejectedMessageProfileId uuid.UUID `json:"rejected_message_profile_id" validate:"required"`
}
