package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/repository"
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
