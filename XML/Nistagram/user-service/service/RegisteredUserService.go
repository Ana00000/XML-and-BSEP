package service

import (
	"../model"
	"../repository"
)

type RegisteredUserService struct {
	Repo * repository.RegisteredUserRepository
}

func (service * RegisteredUserService) CreateRegisteredUser(registeredUser *model.RegisteredUser) error {
	service.Repo.CreateRegisteredUser(registeredUser)
	return nil
}