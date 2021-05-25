package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/repository"
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