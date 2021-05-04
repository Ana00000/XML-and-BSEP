package dto

import (
	"github.com/google/uuid"
)

type AgentRegistrationRequestDTO struct {
	AgentId uuid.UUID `json:"userId"`
}
