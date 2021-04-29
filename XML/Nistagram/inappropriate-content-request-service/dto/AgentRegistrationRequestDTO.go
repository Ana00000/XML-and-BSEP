package dto

import (
	"github.com/google/uuid"
)

type AgentRegistrationRequestDTO struct {
	UserId uuid.UUID `json:"userId"`
}
