package service

import (
	"../model"
	"../repository"
)

type RegisteredUserFollowingsService struct {
	Repo * repository.RegisteredUserFollowingsRepository
}

func (service * RegisteredUserFollowingsService) CreateRegisteredUserFollowings(registeredUserFollowings *model.RegisteredUserFollowings) error {
	service.Repo.CreateRegisteredUserFollowings(registeredUserFollowings)
	return nil
}
