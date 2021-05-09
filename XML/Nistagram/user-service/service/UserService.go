package service

import (
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