package service

import (
	"../model"
	"../repository"
)

type PostAlbumService struct {
	Repo * repository.PostAlbumRepository
}

func (service * PostAlbumService) CreatePostAlbum(postAlbum *model.PostAlbum) error {
	err := service.Repo.CreatePostAlbum(postAlbum)
	if err != nil {
		return err
	}
	return nil
}