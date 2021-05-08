package service

import (
	"../model"
	"../repository"
)

type CampaignChosenGroupService struct {
	Repo * repository.CampaignChosenGroupRepository
}

func (service * CampaignChosenGroupService) CreateCampaignChosenGroup(campaignChosenGroup *model.CampaignChosenGroup) error {
	err := service.Repo.CreateCampaignChosenGroup(campaignChosenGroup)
	if err != nil {
		return err
	}
	return nil
}
