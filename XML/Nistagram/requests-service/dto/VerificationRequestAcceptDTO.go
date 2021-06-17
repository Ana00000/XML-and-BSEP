package dto

import "github.com/google/uuid"

type VerificationRequestAcceptDTO struct {
	ID uuid.UUID `json:"id" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
	RegisteredUserCategory string `json:"registered_user_category" validate:"required"`

}
