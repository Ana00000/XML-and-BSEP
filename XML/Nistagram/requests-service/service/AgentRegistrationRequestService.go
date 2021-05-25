package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
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