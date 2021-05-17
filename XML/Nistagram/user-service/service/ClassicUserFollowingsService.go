package service

import (
	"github.com/google/uuid"
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


func (service * ClassicUserFollowingsService)  CheckIfFollowingUser(classicUserId uuid.UUID, followingUserId uuid.UUID) bool {
	return service.Repo.CheckIfFollowingUser(classicUserId, followingUserId)
}