package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
)

type PostCollectionPostsService struct {
	Repo * repository.PostCollectionPostsRepository
}

func (service * PostCollectionPostsService) CreatePostCollectionPosts(postCollectionPosts *model.PostCollectionPosts) error {
	err := service.Repo.CreatePostCollectionPosts(postCollectionPosts)
	if err != nil {
		return err
	}
	return nil
}
