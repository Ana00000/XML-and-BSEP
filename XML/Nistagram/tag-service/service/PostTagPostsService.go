package service

import (
	"../model"
	"../repository"
)

type PostTagPostsService struct {
	Repo * repository.PostTagPostsRepository
}

func (service * PostTagPostsService) CreatePostTagPosts(postTagPosts *model.PostTagPosts) error {
	service.Repo.CreatePostTagPosts(postTagPosts)
	return nil
}
