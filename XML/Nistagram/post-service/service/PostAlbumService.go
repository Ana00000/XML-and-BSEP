package service

import (
	"../model"
	"../repository"
)

type PostAlbumService struct {
	Repo * repository.PostAlbumRepository
}

func (service * PostAlbumService) CreatePostAlbum(postAlbum *model.PostAlbum) error {
	service.Repo.CreatePostAlbum(postAlbum)
	return nil
}