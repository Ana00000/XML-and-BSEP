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


func (service *VerificationRequestService) FindAllPendingVerificationRequests() []model.VerificationRequest {
	requests := service.Repo.FindAllPendingVerificationRequests()
	if requests != nil {
		return requests
	}
	return nil
}


func (service *VerificationRequestService) UpdateVerificationRequestAccepted(verificationRequestId uuid.UUID) error {
	err := service.Repo.UpdateVerificationRequestAccepted(verificationRequestId)
	if err != nil {
		return err
	}
	return nil
}

func (service *VerificationRequestService) UpdateVerificationRequestRejected(verificationRequestId uuid.UUID) error {
	err := service.Repo.UpdateVerificationRequestRejected(verificationRequestId)
	if err != nil {
		return err
	}
	return nil
}
