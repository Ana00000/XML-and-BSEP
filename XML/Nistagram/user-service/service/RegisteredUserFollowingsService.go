package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
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
