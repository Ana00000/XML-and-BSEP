package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type RegisteredUserService struct {
	Repo * repository.RegisteredUserRepository
}

func (service * RegisteredUserService) CreateRegisteredUser(registeredUser *model.RegisteredUser) error {
	err := service.Repo.CreateRegisteredUser(registeredUser)
	if err != nil {
		return err
	}
	return nil
}

func (service *RegisteredUserService) UpdateRegisteredUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	err := service.Repo.UpdateRegisteredUserConfirmed(userId,isConfirmed)
	if err != nil {
		return err
	}
	return nil
}