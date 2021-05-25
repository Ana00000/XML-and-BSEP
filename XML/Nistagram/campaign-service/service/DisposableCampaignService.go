package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/repository"
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
