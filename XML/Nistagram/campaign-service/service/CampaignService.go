package service

import (
	"../model"
	"../repository"
)

type CampaignService struct {
	Repo * repository.CampaignRepository
}

func (service * CampaignService) CreateCampaign(campaign *model.Campaign) error {
	err := service.Repo.CreateCampaign(campaign)
	if err != nil {
		return err
	}
	return nil
}