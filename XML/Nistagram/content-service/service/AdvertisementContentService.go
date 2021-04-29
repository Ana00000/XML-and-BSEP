package service

import (
	"../model"
	"../repository"
)

type AdvertisementContentService struct {
	Repo * repository.AdvertisementContentRepository
}

func (service * AdvertisementContentService) CreateAdvertisementContent(advertisementContent *model.AdvertisementContent) error {
	service.Repo.CreateAdvertisementContent(advertisementContent)
	return nil
}