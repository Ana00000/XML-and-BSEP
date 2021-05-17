package service

import (
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