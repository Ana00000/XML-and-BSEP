package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/repository"
)

type PostMessageSubstanceService struct {
	Repo *repository.PostMessageSubstanceRepository
}

func (service * PostMessageSubstanceService) CreatePostMessageSubstance(postMessageSubstance *model.PostMessageSubstance) error {
	err := service.Repo.CreatePostMessageSubstance(postMessageSubstance)
	if err != nil {
		return err
	}
	return nil
}


