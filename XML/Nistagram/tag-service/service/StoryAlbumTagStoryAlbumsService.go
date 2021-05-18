package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type StoryAlbumTagStoryAlbumsService struct {
	Repo * repository.StoryAlbumTagStoryAlbumsRepository
}

func (service * StoryAlbumTagStoryAlbumsService) CreateStoryAlbumTagStoryAlbums(storyAlbumTagStoryAlbums *model.StoryAlbumTagStoryAlbums) error {
	err := service.Repo.CreateStoryAlbumTagStoryAlbums(storyAlbumTagStoryAlbums)
	if err != nil {
		return err
	}
	return nil
}
