package service

import (
	"../model"
	"../repository"
)

type PostAlbumContentService struct {
	Repo * repository.PostAlbumContentRepository
}

func (service * PostAlbumContentService) CreatePostAlbumContent(postAlbumContent *model.PostAlbumContent) error {
	service.Repo.CreatePostAlbumContent(postAlbumContent)
	return nil
}
