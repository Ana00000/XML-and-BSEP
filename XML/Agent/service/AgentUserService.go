package service

import (
	"../model"
	"../repository"
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