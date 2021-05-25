package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type ConfirmationTokenService struct {
	Repo * repository.ConfirmationTokenRepository
}

func (service * ConfirmationTokenService) CreateConfirmationToken(confirmationToken *model.ConfirmationToken) error {
	err := service.Repo.CreateConfirmationToken(confirmationToken)
	if err != nil {
		return err
	}
	return nil
}

func (service * ConfirmationTokenService) FindByToken(token uuid.UUID) *model.ConfirmationToken {
	confirmationToken := service.Repo.FindByToken(token)
	return confirmationToken
}

func (service * ConfirmationTokenService) UpdateConfirmationTokenValidity(token uuid.UUID, isValid bool) error {
	err := service.Repo.UpdateConfirmationTokenValidity(token,isValid)
	if err != nil {
		return err
	}
	return nil
}