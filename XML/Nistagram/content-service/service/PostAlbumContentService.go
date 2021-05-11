package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
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
