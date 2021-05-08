package service

import (
	"../model"
	"../repository"
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