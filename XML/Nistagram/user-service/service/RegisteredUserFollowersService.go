package service

import (
	"../model"
	"../repository"
)

type RegisteredUserFollowersService struct {
	Repo * repository.RegisteredUserFollowersRepository
}

func (service * RegisteredUserFollowersService) CreateRegisteredUserFollowers(registeredUserFollowers *model.RegisteredUserFollowers) error {
	service.Repo.CreateRegisteredUserFollowers(registeredUserFollowers)
	return nil
}