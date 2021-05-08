package service

import (
	"../model"
	"../repository"
)

type VerificationRequestService struct {
	Repo * repository.VerificationRequestRepository
}

func (service * VerificationRequestService) CreateVerificationRequest(verificationRequest *model.VerificationRequest) error {
	err := service.Repo.CreateVerificationRequest(verificationRequest)
	if err != nil {
		return err
	}
	return nil
}