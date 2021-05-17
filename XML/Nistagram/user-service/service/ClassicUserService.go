package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type ClassicUserService struct {
	Repo *repository.ClassicUserRepository
}

func (service *ClassicUserService) CreateClassicUser(classicUser *model.ClassicUser) error {
	err := service.Repo.CreateClassicUser(classicUser)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClassicUserService) UpdateClassicUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	err := service.Repo.UpdateClassicUserConfirmed(userId, isConfirmed)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClassicUserService) UpdateClassicUserProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	err := service.Repo.UpdateClassicUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClassicUserService) UpdateClassicUserPassword(userId uuid.UUID, salt string, password string) error {
	err := service.Repo.UpdateClassicUserPassword(userId, salt, password)
	if err != nil {
		return err
	}
	return nil
}

func (service *ClassicUserService) FindSelectedUserById(id uuid.UUID) *dto.SelectedUserDTO {
	user := service.Repo.FindSelectedUserById(id)
	return user
}

func (service *ClassicUserService) FindClassicUserByUserName(userName string) *model.ClassicUser {
	user := service.Repo.FindClassicUserByUserName(userName)
	return user
}
