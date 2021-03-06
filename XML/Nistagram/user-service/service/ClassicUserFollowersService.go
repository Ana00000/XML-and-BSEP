package service

import (
	"github.com/google/uuid"
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

func (service * ClassicUserFollowersService) FindAllFollowersForUser(userId uuid.UUID) []model.ClassicUserFollowers{
	users := service.Repo.FindAllFollowersForUser(userId)
	if users != nil {
		return users
	}
	return nil
}

func (service * ClassicUserFollowersService) CheckIfFollowers(classicUserId uuid.UUID, followerUserId uuid.UUID, ) bool {
	return service.Repo.CheckIfFollowers(classicUserId, followerUserId)
}

func (service * ClassicUserFollowersService) FindAllMutualFollowerForUser(userId uuid.UUID) []model.ClassicUserFollowers{
	users := service.Repo.FindAllMutualFollowerForUser(userId)
	if users != nil {
		return users
	}
	return nil
}

func (service ClassicUserFollowersService) FindFollowerByUsersIDs(followingUserID uuid.UUID,classicUserID uuid.UUID) *model.ClassicUserFollowers {
	return service.Repo.FindFollowerByUsersIDs(followingUserID,classicUserID)
}

func (service *ClassicUserFollowersService) RemoveClassicUserFollower(id uuid.UUID) {
	service.Repo.RemoveClassicUserFollower(id)
}