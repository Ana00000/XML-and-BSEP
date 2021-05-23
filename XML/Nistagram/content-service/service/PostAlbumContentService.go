package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
	postsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
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

func (service *PostAlbumContentService) FindAllContentsForPostAlbums(allPostAlbums []postsModel.PostAlbum) []model.PostAlbumContent {
	postAlbums := service.Repo.FindAllContentsForPostAlbums(allPostAlbums)
	if postAlbums != nil {
		return postAlbums
	}
	return nil
}

func (service *PostAlbumContentService) FindAllContentsForPostAlbum(album *postsModel.PostAlbum) []model.PostAlbumContent {
	contents := service.Repo.FindAllContentsForPostAlbum(album)
	if contents != nil {
		return contents
	}
	return nil
}
