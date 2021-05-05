package service

import (
	"../model"
	"../repository"
)

type CampaignChosenGroupService struct {
	Repo * repository.CampaignChosenGroupRepository
}

func (service * CampaignChosenGroupService) CreateCampaignChosenGroup(campaignChosenGroup *model.CampaignChosenGroup) error {
	service.Repo.CreateCampaignChosenGroup(campaignChosenGroup)
	return nil
}
