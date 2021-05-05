package service

import (
	"../model"
	"../repository"
)

type StoryAlbumContentService struct {
	Repo * repository.StoryAlbumContentRepository
}

func (service * StoryAlbumContentService) CreateStoryAlbumContent(storyAlbumContent *model.StoryAlbumContent) error {
	service.Repo.CreateStoryAlbumContent(storyAlbumContent)
	return nil
}
