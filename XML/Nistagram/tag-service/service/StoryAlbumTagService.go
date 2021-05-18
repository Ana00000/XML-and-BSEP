package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type StoryAlbumTagService struct {
	Repo * repository.StoryAlbumTagRepository
}

func (service * StoryAlbumTagService) CreateStoryAlbumTag(storyAlbumTag *model.StoryAlbumTag) error {
	err := service.Repo.CreateStoryAlbumTag(storyAlbumTag)
	if err != nil {
		return err
	}
	return nil
}