package service

import (
	"../model"
	"../repository"
)

type PostCollectionPostsService struct {
	Repo * repository.PostCollectionPostsRepository
}

func (service * PostCollectionPostsService) CreatePostCollectionPosts(postCollectionPosts *model.PostCollectionPosts) error {
	service.Repo.CreatePostCollectionPosts(postCollectionPosts)
	return nil
}
