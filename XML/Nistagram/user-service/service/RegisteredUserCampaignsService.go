package service

import (
	"../model"
	"../repository"
)

type RegisteredUserCampaignsService struct {
	Repo * repository.RegisteredUserCampaignsRepository
}

func (service * RegisteredUserCampaignsService) CreateRegisteredUserCampaigns(registeredUserCampaigns *model.RegisteredUserCampaigns) error {
	err := service.Repo.CreateRegisteredUserCampaigns(registeredUserCampaigns)
	if err != nil {
		return err
	}
	return nil
}
