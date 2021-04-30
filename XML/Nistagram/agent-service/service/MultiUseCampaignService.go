package service

import (
	"../model"
	"../repository"
)

type MultiUseCampaignService struct {
	Repo * repository.MultiUseCampaignRepository
}

func (service * MultiUseCampaignService) CreateMultiUseCampaign(multiUseCampaign *model.MultiUseCampaign) error {
	service.Repo.CreateMultiUseCampaign(multiUseCampaign)
	return nil
}