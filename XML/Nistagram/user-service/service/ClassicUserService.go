package service

import (
	"../model"
	"../repository"
)

type ClassicUserService struct {
	Repo * repository.ClassicUserRepository
}

func (service * ClassicUserService) CreateClassicUser(classicUser *model.ClassicUser) error {
	err := service.Repo.CreateClassicUser(classicUser)
	if err != nil {
		return err
	}
	return nil
}