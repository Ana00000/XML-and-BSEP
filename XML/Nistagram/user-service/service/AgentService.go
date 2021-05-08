package service

import (
	"../model"
	"../repository"
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