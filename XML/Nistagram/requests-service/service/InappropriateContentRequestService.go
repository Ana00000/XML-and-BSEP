package service

import (
	"../model"
	"../repository"
)

type InappropriateContentRequestService struct {
	Repo * repository.InappropriateContentRequestRepository
}

func (service * InappropriateContentRequestService) CreateInappropriateContentRequest(inappropriateContentRequest *model.InappropriateContentRequest) error {
	service.Repo.CreateInappropriateContentRequest(inappropriateContentRequest)
	return nil
}