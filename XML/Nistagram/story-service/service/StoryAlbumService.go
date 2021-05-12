package service

import (
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