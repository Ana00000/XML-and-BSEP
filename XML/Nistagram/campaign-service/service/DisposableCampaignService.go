package service

import (
	"../model"
	"../repository"
)

type DisposableCampaignService struct {
	Repo * repository.DisposableCampaignRepository
}

func (service * DisposableCampaignService) CreateDisposableCampaign(disposableCampaign *model.DisposableCampaign) error {
	service.Repo.CreateDisposableCampaign(disposableCampaign)
	return nil
}
