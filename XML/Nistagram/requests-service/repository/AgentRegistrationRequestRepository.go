package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type AgentRegistrationRequestRepository struct {
	Database * gorm.DB
}

func (repo * AgentRegistrationRequestRepository) CreateAgentRegistrationRequest(agentRegistrationRequest *model.AgentRegistrationRequest) error {
	result := repo.Database.Create(agentRegistrationRequest)
	fmt.Print(result)
	return nil
}