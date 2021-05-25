package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
)

type AdvertisementContentService struct {
	Repo * repository.AdvertisementContentRepository
}

func (service * AdvertisementContentService) CreateAdvertisementContent(advertisementContent *model.AdvertisementContent) error {
	err := service.Repo.CreateAdvertisementContent(advertisementContent)
	if err != nil {
		return err
	}
	return nil
}