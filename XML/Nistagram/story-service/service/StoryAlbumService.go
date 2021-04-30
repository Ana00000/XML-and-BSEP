package service

import (
	"../model"
	"../repository"
)

type StoryAlbumService struct {
	Repo * repository.StoryAlbumRepository
}

func (service * StoryAlbumService) CreateStoryAlbum(storyAlbum *model.StoryAlbum) error {
	service.Repo.CreateStoryAlbum(storyAlbum)
	return nil
}