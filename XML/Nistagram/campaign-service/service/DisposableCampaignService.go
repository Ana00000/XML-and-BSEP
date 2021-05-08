package service

import (
	"../model"
	"../repository"
)

type DisposableCampaignService struct {
	Repo * repository.DisposableCampaignRepository
}

func (service * DisposableCampaignService) CreateDisposableCampaign(disposableCampaign *model.DisposableCampaign) error {
	err := service.Repo.CreateDisposableCampaign(disposableCampaign)
	if err != nil {
		return err
	}
	return nil
}
