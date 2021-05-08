package service

import (
	"../model"
	"../repository"
)

type AgentRegistrationRequestService struct {
	Repo * repository.AgentRegistrationRequestRepository
}

func (service * AgentRegistrationRequestService) CreateAgentRegistrationRequest(agentRegistrationRequest *model.AgentRegistrationRequest) error {
	err := service.Repo.CreateAgentRegistrationRequest(agentRegistrationRequest)
	if err != nil {
		return err
	}
	return nil
}