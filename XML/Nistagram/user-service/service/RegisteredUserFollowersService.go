package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
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