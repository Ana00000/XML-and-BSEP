package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
)

type VerificationRequestService struct {
	Repo *repository.VerificationRequestRepository
}

func (service *VerificationRequestService) CreateVerificationRequest(verificationRequest *model.VerificationRequest) error {
	err := service.Repo.CreateVerificationRequest(verificationRequest)
	if err != nil {
		return err
	}
	return nil
}

func (service *VerificationRequestService) FindById(id uuid.UUID) *model.VerificationRequest {
	request := service.Repo.FindById(id)
	if request != nil {
		return request
	}
	return nil
}

/*
func (service *VerificationRequestService) FindAllPendingFollowerRequestsForUser(userId uuid.UUID) []model.FollowRequest {
	requests := service.Repo.FindAllPendingFollowRequestsForUser(userId)
	if requests != nil {
		return requests
	}
	return nil
}

func (service *VerificationRequestService) FindFollowRequest(classicUserId uuid.UUID, followerUserId uuid.UUID) *model.FollowRequest {
	request := service.Repo.FindFollowRequest(classicUserId, followerUserId)
	if request != nil {
		return request
	}
	return nil
}

func (service *VerificationRequestService) UpdateFollowRequestPending(followRequestId uuid.UUID) error {
	err := service.Repo.UpdateFollowRequestPending(followRequestId)
	if err != nil {
		return err
	}
	return nil
}

func (service *VerificationRequestService) UpdateFollowRequestAccepted(followRequestId uuid.UUID) error {
	err := service.Repo.UpdateFollowRequestAccepted(followRequestId)
	if err != nil {
		return err
	}
	return nil
}

func (service *VerificationRequestService) UpdateFollowRequestRejected(followRequestId uuid.UUID) error {
	err := service.Repo.UpdateFollowRequestRejected(followRequestId)
	if err != nil {
		return err
	}
	return nil
}
*/