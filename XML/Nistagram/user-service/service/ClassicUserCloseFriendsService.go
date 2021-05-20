package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type ClassicUserCloseFriendsService struct {
	Repo * repository.ClassicUserCloseFriendsRepository
}

func (service * ClassicUserCloseFriendsService) CreateClassicUserCloseFriends(classicUserCloseFriends *model.ClassicUserCloseFriends) error {
	err := service.Repo.CreateClassicUserCloseFriends(classicUserCloseFriends)
	if err != nil {
		return err
	}
	return nil
}
