package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type PostAlbumTagPostAlbumsService struct {
	Repo * repository.PostAlbumTagPostAlbumsRepository
}

func (service * PostAlbumTagPostAlbumsService) CreatePostAlbumTagPostAlbums(postAlbumTagPostAlbums *model.PostAlbumTagPostAlbums) error {
	err := service.Repo.CreatePostAlbumTagPostAlbums(postAlbumTagPostAlbums)
	if err != nil {
		return err
	}
	return nil
}
