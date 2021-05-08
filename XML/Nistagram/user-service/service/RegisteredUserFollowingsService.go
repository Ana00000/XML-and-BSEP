package service

import (
	"../model"
	"../repository"
)

type RegisteredUserFollowingsService struct {
	Repo * repository.RegisteredUserFollowingsRepository
}

func (service * RegisteredUserFollowingsService) CreateRegisteredUserFollowings(registeredUserFollowings *model.RegisteredUserFollowings) error {
	err := service.Repo.CreateRegisteredUserFollowings(registeredUserFollowings)
	if err != nil {
		return err
	}
	return nil
}
