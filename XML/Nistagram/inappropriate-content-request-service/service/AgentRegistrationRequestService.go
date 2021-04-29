package service

import (
	"../model"
	"../repository"
)

type AgentRegistrationRequestService struct {
	Repo * repository.AgentRegistrationRequestRepository
}

func (service * AgentRegistrationRequestService) CreateAgentRegistrationRequest(agentRegistrationRequest *model.AgentRegistrationRequest) error {
	service.Repo.CreateAgentRegistrationRequest(agentRegistrationRequest)
	return nil
}