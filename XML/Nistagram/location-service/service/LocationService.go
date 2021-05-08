package service

import (
	"../model"
	"../repository"
)

type LocationService struct {
	Repo * repository.LocationRepository
}

func (service * LocationService) CreateLocation(location *model.Location) error {
	err := service.Repo.CreateLocation(location)
	if err != nil {
		return err
	}
	return nil
}
