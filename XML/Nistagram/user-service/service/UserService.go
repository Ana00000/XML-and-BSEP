package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type UserService struct {
	Repo * repository.UserRepository
}

func (service * UserService) CreateUser(user *model.User) error {
	err := service.Repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (service * UserService) FindAllUsers() []model.User{
	users := service.Repo.FindAllUsers()
	if users != nil {
		return users
	}
	return nil
}

func (service *UserService) FindByUserName(userName string) *model.User {
	user := service.Repo.FindByUserName(userName)
	return user
}

func (service *UserService) FindByEmail(email string) *model.User {
	user := service.Repo.FindByEmail(email)
	return user
}

func (service *UserService) FindByID(ID uuid.UUID) *model.User {
	user := service.Repo.FindByID(ID)
	return user
}
func (service *UserService) UpdateUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	err := service.Repo.UpdateUserConfirmed(userId,isConfirmed)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateUserProfileInfo(user *dto.UserUpdateProfileInfoDTO) error {
	err := service.Repo.UpdateUserProfileInfo(user)
}

func (service *UserService) UpdateUserPassword(userId uuid.UUID, password string) error {
	err := service.Repo.UpdateUserPassword(userId,password)
	if err != nil {
		return err
	}
	return nil
}