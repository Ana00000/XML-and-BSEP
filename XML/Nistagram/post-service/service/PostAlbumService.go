package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
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