package service

import (
	"../model"
	"../repository"
)

type PostAlbumContentService struct {
	Repo * repository.PostAlbumContentRepository
}

func (service * PostAlbumContentService) CreatePostAlbumContent(postAlbumContent *model.PostAlbumContent) error {
	err := service.Repo.CreatePostAlbumContent(postAlbumContent)
	if err != nil {
		return err
	}
	return nil
}
