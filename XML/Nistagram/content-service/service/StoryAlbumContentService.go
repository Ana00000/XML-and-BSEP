package service

import (
	"../model"
	"../repository"
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
