package model

import "github.com/google/uuid"

type Agent struct {
	ClassicUser
	AgentRegistrationRequestId uuid.UUID `json:"agent_registration_request_id" gorm:"not null"`
}
