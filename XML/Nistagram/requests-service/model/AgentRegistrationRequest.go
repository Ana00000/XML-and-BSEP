package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AgentRegistrationRequest struct {
	ID uuid.UUID `json:"id"`
	AgentId uuid.UUID `json:"agent" gorm:"not null"`
}

func(agentRegistrationRequest * AgentRegistrationRequest) BeforeCreate(scope *gorm.DB) error {
	agentRegistrationRequest.ID = uuid.New()
	return nil
}
