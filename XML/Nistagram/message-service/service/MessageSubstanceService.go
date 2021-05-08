package service

import (
	"../model"
	"../repository"
)

type MessageSubstanceService struct {
	Repo *repository.MessageSubstanceRepository
}

func (service * MessageSubstanceService) CreateMessageSubstance(messageSubstance *model.MessageSubstance) error {
	err := service.Repo.CreateMessageSubstance(messageSubstance)
	if err != nil {
		return err
	}
	return nil
}
