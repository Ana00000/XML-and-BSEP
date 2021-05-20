package service

import (
	"github.com/google/uuid"
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

func (service * ClassicUserCloseFriendsService)  CheckIfCloseFriend(classicUserId uuid.UUID, closeFriendUserId uuid.UUID) bool {
	return service.Repo.CheckIfCloseFriend(classicUserId, closeFriendUserId)
}

