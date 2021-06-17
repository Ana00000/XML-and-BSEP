package dto

import "github.com/google/uuid"

type MuteUserDTO struct {
	LoggedInUser uuid.UUID `json:"logged_in_user"`
	MutedUser uuid.UUID `json:"muted_user"`
}