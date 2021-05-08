package service

import (
	"../model"
	"../repository"
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