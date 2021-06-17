package dto

import "github.com/google/uuid"

//BlockUserDTO

type BlockUserDTO struct {
	LoggedInUser uuid.UUID `json:"logged_in_user"`
	BlockedUser uuid.UUID `json:"blocked_user"`
}