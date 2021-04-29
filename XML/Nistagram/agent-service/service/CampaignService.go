package service

import (
	"../model"
	"../repository"
)

type CampaignService struct {
	Repo * repository.CampaignRepository
}

func (service * CampaignService) CreateCampaign(campaign *model.Campaign) error {
	service.Repo.CreateCampaign(campaign)
	return nil
}