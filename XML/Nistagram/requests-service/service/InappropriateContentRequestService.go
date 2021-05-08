package service

import (
	"../model"
	"../repository"
)

type InappropriateContentRequestService struct {
	Repo * repository.InappropriateContentRequestRepository
}

func (service * InappropriateContentRequestService) CreateInappropriateContentRequest(inappropriateContentRequest *model.InappropriateContentRequest) error {
	err := service.Repo.CreateInappropriateContentRequest(inappropriateContentRequest)
	if err != nil {
		return err
	}
	return nil
}