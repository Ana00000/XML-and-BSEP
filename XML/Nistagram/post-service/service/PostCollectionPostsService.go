package service

import (
	"github.com/google/uuid"
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

func (service * PostCollectionPostsService) FindAllPostCollectionPostsForPost(postId uuid.UUID) []model.PostCollectionPosts{
	postCollectionPosts := service.Repo.FindAllPostCollectionPostsForPost(postId)
	if postCollectionPosts != nil {
		return postCollectionPosts
	}
	return nil
}