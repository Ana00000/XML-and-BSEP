package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
)

type ClassicUserCampaignsService struct {
	Repo * repository.ClassicUserCampaignsRepository
}

func (service * ClassicUserCampaignsService) CreateClassicUserCampaigns(classicUserCampaigns *model.ClassicUserCampaigns) error {
	err := service.Repo.CreateClassicUserCampaigns(classicUserCampaigns)
	if err != nil {
		return err
	}
	return nil
}
