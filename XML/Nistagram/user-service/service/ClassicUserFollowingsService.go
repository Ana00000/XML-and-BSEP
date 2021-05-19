package service

import (
	"github.com/google/uuid"
	requestModel "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
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


func (service * ClassicUserFollowingsService)  CheckFollowingStatus(classicUserId uuid.UUID, followingUserId uuid.UUID, followRequests []requestModel.FollowRequest) string {
	return service.Repo.CheckFollowingStatus(classicUserId, followingUserId, followRequests)
}