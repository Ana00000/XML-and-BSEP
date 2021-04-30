package service

import (
	"../model"
	"../repository"
)

type PostService struct {
	Repo * repository.PostRepository
}

func (service * PostService) CreatePost(post *model.Post) error {
	service.Repo.CreatePost(post)
	return nil
}
