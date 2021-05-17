package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type PostAlbumTagService struct {
	Repo * repository.PostAlbumTagRepository
}

func (service * PostAlbumTagService) CreatePostAlbumTag(postAlbumTag *model.PostAlbumTag) error {
	err := service.Repo.CreatePostAlbumTag(postAlbumTag)
	if err != nil {
		return err
	}
	return nil
}