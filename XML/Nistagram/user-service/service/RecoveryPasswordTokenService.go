package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type RecoveryPasswordTokenService struct {
	Repo * repository.RecoveryPasswordTokenRepository
}

func (service * RecoveryPasswordTokenService) CreateRecoveryPasswordToken(recoveryPasswordToken *model.RecoveryPasswordToken) error {
	err := service.Repo.CreateRecoveryPasswordToken(recoveryPasswordToken)
	if err != nil {
		return err
	}
	return nil
}

func (service * RecoveryPasswordTokenService) FindByToken(token uuid.UUID) *model.RecoveryPasswordToken {
	recoveryPasswordToken := service.Repo.FindByToken(token)
	return recoveryPasswordToken
}

func (service * RecoveryPasswordTokenService) UpdateRecoveryPasswordTokenValidity(token uuid.UUID, isValid bool) error {
	err := service.Repo.UpdateRecoveryPasswordTokenValidity(token,isValid)
	if err != nil {
		return err
	}
	return nil
}