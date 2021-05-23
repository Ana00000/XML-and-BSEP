package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/repository"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	storyModel "github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
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

func (service *LocationService) FindAllLocationsForPosts(allPosts []postsModel.SinglePost) []model.Location {
	locations := service.Repo.FindAllLocationsForPosts(allPosts)
	if locations != nil {
		return locations
	}
	return nil
}

func (service *LocationService) FindAllLocationsForPost(post *postsModel.SinglePost) []model.Location {
	locations := service.Repo.FindAllLocationsForPost(post)
	if locations != nil {
		return locations
	}
	return nil
}

func (service *LocationService) FindAllLocationsForStories(allStories []storyModel.SingleStory) []model.Location {
	locations := service.Repo.FindAllLocationsForStories(allStories)
	if locations != nil {
		return locations
	}
	return nil
}

func (service *LocationService) FindAllLocationsForStory(story *storyModel.SingleStory) []model.Location {
	locations := service.Repo.FindAllLocationsForStory(story)
	if locations != nil {
		return locations
	}
	return nil
}

func (service *LocationService) FindLocationIdByLocationString(locationString string) model.Location {
	return service.Repo.FindLocationIdByLocationString(locationString)
}

func (service *LocationService) FindAllLocationsForPostAlbums(albums []postsModel.PostAlbum) []model.Location {
	locations := service.Repo.FindAllLocationsForPostAlbums(albums)
	if locations != nil {
		return locations
	}
	return nil
}

func (service *LocationService) FindAllLocationsForPostAlbum(album *postsModel.PostAlbum) []model.Location {
	locations := service.Repo.FindAllLocationsForPostAlbum(album)
	if locations != nil {
		return locations
	}
	return nil
}