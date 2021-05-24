package service

import (
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
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

func (service *PostAlbumTagPostAlbumsService) FindAllTagsForPostAlbumTagPostAlbums(albums []postsModel.PostAlbum) []model.PostAlbumTagPostAlbums {
	tags := service.Repo.FindAllTagsForPostAlbumTagPostAlbums(albums)
	if tags != nil {
		return tags
	}
	return nil
}

func (service *PostAlbumTagPostAlbumsService) FindAllTagsForPostAlbum(album *postsModel.PostAlbum) []model.PostAlbumTagPostAlbums {
	tags := service.Repo.FindAllTagsForPostAlbum(album)
	if tags != nil {
		return tags
	}
	return nil
}