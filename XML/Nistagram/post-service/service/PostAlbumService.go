package service

import (
	"github.com/google/uuid"
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

func (service *PostAlbumService) FindAllAlbumPostsForUser(ID uuid.UUID) []model.PostAlbum {
	albumPosts := service.Repo.FindAllAlbumPostsForUser(ID)
	if albumPosts != nil {
		return albumPosts
	}
	return nil
}

func (service *PostAlbumService) FindByID(ID uuid.UUID) *model.PostAlbum {
	postAlbum := service.Repo.FindByID(ID)
	return postAlbum
}