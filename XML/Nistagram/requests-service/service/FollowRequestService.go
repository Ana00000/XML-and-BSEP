package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
)

type FollowRequestService struct {
	Repo * repository.FollowRequestRepository
}

func (service * FollowRequestService) CreateFollowRequest(followRequest *model.FollowRequest) error {
	err := service.Repo.CreateFollowRequest(followRequest)
	if err != nil {
		return err
	}
	return nil
}

func (service * FollowRequestService) FindById(id uuid.UUID) *model.FollowRequest{
	requests := service.Repo.FindById(id)
	if requests != nil {
		return requests
	}
	return nil
}

func (service * FollowRequestService) FindAllFollowerRequestsForUser(userId uuid.UUID) []model.FollowRequest{
	requests := service.Repo.FindAllFollowRequestsForUser(userId)
	if requests != nil {
		return requests
	}
	return nil
}
