package service

import (
	"../model"
	"../repository"
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