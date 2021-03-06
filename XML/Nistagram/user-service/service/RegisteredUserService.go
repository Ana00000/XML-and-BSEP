package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type RegisteredUserService struct {
	Repo *repository.RegisteredUserRepository
}

func (service *RegisteredUserService) CreateRegisteredUser(registeredUser *model.RegisteredUser) error {
	err := service.Repo.CreateRegisteredUser(registeredUser)
	if err != nil {
		return err
	}
	return nil
}

func (service *RegisteredUserService) UpdateRegisteredUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	err := service.Repo.UpdateRegisteredUserConfirmed(userId, isConfirmed)
	if err != nil {
		return err
	}
	return nil
}

func (service *RegisteredUserService) UpdateRegisteredUserProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	err := service.Repo.UpdateRegisteredUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *RegisteredUserService) UpdateRegisteredUserPassword(userId uuid.UUID, salt string, password string) error {
	err := service.Repo.UpdateRegisteredUserPassword(userId, salt, password)
	if err != nil {
		return err
	}
	return nil
}

func (service *RegisteredUserService) UpdateUserCategory(userId uuid.UUID, category model.RegisteredUserCategory) error {
	err := service.Repo.UpdateUserCategory(userId, category)
	if err != nil {
		return err
	}
	return nil
}

func (service *RegisteredUserService) UpdateOfficialDocumentPath(id uuid.UUID, path string) error{
	err := service.Repo.UpdateOfficialDocumentPath(id, path)
	if err != nil {
		return err
	}
	return nil
}