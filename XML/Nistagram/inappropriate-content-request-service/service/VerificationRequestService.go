package service

import (
	"../model"
	"../repository"
)

type VerificationRequestService struct {
	Repo * repository.VerificationRequestRepository
}

func (service * VerificationRequestService) CreateVerificationRequest(verificationRequest *model.VerificationRequest) error {
	service.Repo.CreateVerificationRequest(verificationRequest)
	return nil
}