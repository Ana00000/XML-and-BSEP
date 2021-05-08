package service

import (
	"../model"
	"../repository"
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