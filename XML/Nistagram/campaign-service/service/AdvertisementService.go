package service

import (
	"../model"
	"../repository"
)

type AdvertisementService struct {
	Repo * repository.AdvertisementRepository
}

func (service * AdvertisementService) CreateAdvertisement(advertisement *model.Advertisement) error {
	service.Repo.CreateAdvertisement(advertisement)
	return nil
}