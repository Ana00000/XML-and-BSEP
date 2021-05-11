package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type AgentService struct {
	Repo * repository.AgentRepository
}

func (service * AgentService) CreateAgent(agent *model.Agent) error {
	err := service.Repo.CreateAgent(agent)
	if err != nil {
		return err
	}
	return nil
}

func (service *AgentService) UpdateAgentProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	err := service.Repo.UpdateAgentProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}