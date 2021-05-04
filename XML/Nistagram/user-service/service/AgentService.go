package service

import (
	"../model"
	"../repository"
)

type AgentService struct {
	Repo * repository.AgentRepository
}

func (service * AgentService) CreateAgent(agent *model.Agent) error {
	service.Repo.CreateAgent(agent)
	return nil
}