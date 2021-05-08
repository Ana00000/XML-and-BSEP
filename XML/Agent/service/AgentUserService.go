package service

import (
	"../model"
	"../repository"
)

type AgentUserService struct {
	Repo * repository.AgentUserRepository
}

func (service * AgentUserService) CreateUser(user *model.AgentUser) error {
	service.Repo.CreateAgentUser(user)
	return nil
}