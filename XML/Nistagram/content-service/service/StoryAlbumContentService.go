package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
)

type StoryAlbumContentService struct {
	Repo * repository.StoryAlbumContentRepository
}

func (service * StoryAlbumContentService) CreateStoryAlbumContent(storyAlbumContent *model.StoryAlbumContent) error {
	err := service.Repo.CreateStoryAlbumContent(storyAlbumContent)
	if err != nil {
		return err
	}
	return nil
}
