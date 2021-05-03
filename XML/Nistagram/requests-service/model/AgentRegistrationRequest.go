package model

import (
	agentPath "../../user-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AgentRegistrationRequest struct {
	ID uuid.UUID `json:"id"`
	Agent agentPath.Agent `json:"agent" gorm:"foreignKey:AgentRegistrationRequestId"`
}

func(agentRegistrationRequest * AgentRegistrationRequest) BeforeCreate(scope *gorm.DB) error {
	agentRegistrationRequest.ID = uuid.New()
	return nil
}
