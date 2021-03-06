package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type ClassicUserFollowingsService struct {
	Repo * repository.ClassicUserFollowingsRepository
}

func (service * ClassicUserFollowingsService) CreateClassicUserFollowings(classicUserFollowings *model.ClassicUserFollowings) error {
	err := service.Repo.CreateClassicUserFollowings(classicUserFollowings)
	if err != nil {
		return err
	}
	return nil
}


func (service * ClassicUserFollowingsService)  CheckFollowingStatus(classicUserId uuid.UUID, followingUserId uuid.UUID, followRequests []dto.FollowRequestForUserDTO) string {
	return service.Repo.CheckFollowingStatus(classicUserId, followingUserId, followRequests)
}

func (service *ClassicUserFollowingsService) FindAllValidFollowingsForUser(ID uuid.UUID, allValidUsers []model.ClassicUser) []model.ClassicUserFollowings {
	followings := service.Repo.FindAllValidFollowingsForUser(ID, allValidUsers)
	if followings != nil {
		return followings
	}
	return nil
}

func (service * ClassicUserFollowingsService) CheckIfFollowingPostStory(followingUserId uuid.UUID, classicUserId uuid.UUID) bool {
	return service.Repo.CheckIfFollowingPostStory(followingUserId, classicUserId)

}

func (service *ClassicUserFollowingsService) FindAllUserWhoFollowUserId(ID uuid.UUID, allValidUsers []model.ClassicUser) []model.ClassicUserFollowings {
	followings := service.Repo.FindAllUserWhoFollowUserId(ID, allValidUsers)
	if followings != nil {
		return followings
	}
	return nil
}

func (service *ClassicUserFollowingsService) FindAllUsersWhoFollowUserId(ID uuid.UUID, allValidUsers []model.ClassicUser) []model.ClassicUser {
	followings := service.Repo.FindAllUsersWhoFollowUserId(ID, allValidUsers)
	if followings != nil {
		return followings
	}
	return nil
}

func (service *ClassicUserFollowingsService) FindFollowingByUsersIDs(followingUserID uuid.UUID,classicUserID uuid.UUID) *model.ClassicUserFollowings {
	return service.Repo.FindFollowingByUsersIDs(followingUserID,classicUserID)
}

func (service *ClassicUserFollowingsService) RemoveClassicUserFollowing(id uuid.UUID) {
	service.Repo.RemoveClassicUserFollowing(id)
}