package service

import (
	"../model"
	"../repository"
)

type ClassicUserService struct {
	Repo * repository.ClassicUserRepository
}

func (service * ClassicUserService) CreateClassicUser(classicUser *model.ClassicUser) error {
	service.Repo.CreateClassicUser(classicUser)
	return nil
}