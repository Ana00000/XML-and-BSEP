package service

import (
	"../model"
	"../repository"
)

type RegisteredUserCampaignsService struct {
	Repo * repository.RegisteredUserCampaignsRepository
}

func (service * RegisteredUserCampaignsService) CreateRegisteredUserCampaigns(registeredUserCampaigns *model.RegisteredUserCampaigns) error {
	service.Repo.CreateRegisteredUserCampaigns(registeredUserCampaigns)
	return nil
}
