package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type ClassicUserFollowersService struct {
	Repo * repository.ClassicUserFollowersRepository
}

func (service * ClassicUserFollowersService) CreateClassicUserFollowers(classicUserFollowers *model.ClassicUserFollowers) error {
	err := service.Repo.CreateClassicUserFollowers(classicUserFollowers)
	if err != nil {
		return err
	}
	return nil
}