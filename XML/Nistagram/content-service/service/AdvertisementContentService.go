package service

import (
	"../model"
	"../repository"
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