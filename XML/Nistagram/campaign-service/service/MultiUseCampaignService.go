package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/repository"
)

type MultiUseCampaignService struct {
	Repo * repository.MultiUseCampaignRepository
}

func (service * MultiUseCampaignService) CreateMultiUseCampaign(multiUseCampaign *model.MultiUseCampaign) error {
	err := service.Repo.CreateMultiUseCampaign(multiUseCampaign)
	if err != nil {
		return err
	}
	return nil
}