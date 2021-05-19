package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/repository"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
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

func (service *LocationService) FindByID(ID uuid.UUID) *model.Location {
	location := service.Repo.FindByID(ID)
	return location
}

func (service *LocationService) FindAllLocationsForPosts(allPosts []postsModel.Post) []model.Location {
	locations := service.Repo.FindAllLocationsForPosts(allPosts)
	if locations != nil {
		return locations
	}
	return nil
}
