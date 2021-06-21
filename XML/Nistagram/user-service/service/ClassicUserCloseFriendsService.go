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

func (service * ClassicUserCloseFriendsService) FindAllCloseFriendsForUser(userId uuid.UUID) []model.ClassicUserCloseFriends{
	users := service.Repo.FindAllCloseFriendsForUser(userId)
	if users != nil {
		return users
	}
	return nil
}

func (service ClassicUserCloseFriendsService) FindCloseFriendByUsersIDs(closeFriendID uuid.UUID,classicUserID uuid.UUID) *model.ClassicUserCloseFriends {
	return service.Repo.FindCloseFriendByUsersIDs(closeFriendID,classicUserID)
}

func (service *ClassicUserCloseFriendsService) RemoveClassicUserCloseFriend(id uuid.UUID) {
	service.Repo.RemoveClassicUserCloseFriend(id)
}
