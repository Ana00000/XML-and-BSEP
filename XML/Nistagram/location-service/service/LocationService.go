package service

import (
	"../model"
	"../repository"
)

type LocationService struct {
	Repo * repository.LocationRepository
}

func (service * LocationService) CreateLocation(location *model.Location) error {
	service.Repo.CreateLocation(location)
	return nil
}
