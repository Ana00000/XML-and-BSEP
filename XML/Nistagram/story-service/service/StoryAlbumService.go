package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/repository"
)

type StoryAlbumService struct {
	Repo * repository.StoryAlbumRepository
}

func (service * StoryAlbumService) CreateStoryAlbum(storyAlbum *model.StoryAlbum) error {
	err := service.Repo.CreateStoryAlbum(storyAlbum)
	if err != nil {
		return err
	}
	return nil
}

func (service *StoryAlbumService) FindAllAlbumStoriesForUser(ID uuid.UUID) []model.StoryAlbum {
	albumStories := service.Repo.FindAllAlbumStoriesForUser(ID)
	if albumStories != nil {
		return albumStories
	}
	return nil
}