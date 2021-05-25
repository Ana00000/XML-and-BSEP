package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/repository"
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
