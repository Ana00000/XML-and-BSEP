package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/repository"
)

type AdvertisementService struct {
	Repo * repository.AdvertisementRepository
}

func (service * AdvertisementService) CreateAdvertisement(advertisement *model.Advertisement) error {
	err := service.Repo.CreateAdvertisement(advertisement)
	if err != nil {
		return err
	}
	return nil
}