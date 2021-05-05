package service

import (
	"../model"
	"../repository"
)

type MessageSubstanceService struct {
	Repo *repository.MessageSubstanceRepository
}

func (service * MessageSubstanceService) CreateMessageSubstance(messageSubstance *model.MessageSubstance) error {
	service.Repo.CreateMessageSubstance(messageSubstance)
	return nil
}
