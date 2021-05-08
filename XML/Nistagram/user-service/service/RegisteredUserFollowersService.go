package service

import (
	"../model"
	"../repository"
)

type RegisteredUserFollowersService struct {
	Repo * repository.RegisteredUserFollowersRepository
}

func (service * RegisteredUserFollowersService) CreateRegisteredUserFollowers(registeredUserFollowers *model.RegisteredUserFollowers) error {
	err := service.Repo.CreateRegisteredUserFollowers(registeredUserFollowers)
	if err != nil {
		return err
	}
	return nil
}