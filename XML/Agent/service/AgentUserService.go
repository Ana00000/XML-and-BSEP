package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Agent/model"
	"github.com/xml/XML-and-BSEP/XML/Agent/repository"
)

type AgentUserService struct {
	Repo * repository.AgentUserRepository
}

func (service * AgentUserService) CreateAgentUser(user *model.AgentUser) error {
	err := service.Repo.CreateAgentUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (service * AgentUserService) FindAgentByUserName(userName string) *model.AgentUser {
	agent := service.Repo.FindAgentByUserName(userName)
	return agent
}

func (service * AgentUserService) FindAgentByEmail(email string) *model.AgentUser {
	agent := service.Repo.FindAgentByEmail(email)
	return agent
}

func (service * AgentUserService) FindAgentByID(ID uuid.UUID) *model.AgentUser {
	agent := service.Repo.FindAgentByID(ID)
	return agent
}