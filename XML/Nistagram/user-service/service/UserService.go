package service

import (
	"../model"
	"../repository"
)

type UserService struct {
	Repo * repository.UserRepository
}

func (service * UserService) CreateUser(user *model.User) error {
	service.Repo.CreateUser(user)
	return nil
}