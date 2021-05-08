package service

import (
	"../model"
	"../repository"
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
