package service

import (
	"../model"
	"../repository"
)

type PostService struct {
	Repo * repository.PostRepository
}

func (service * PostService) CreatePost(post *model.Post) error {
	err := service.Repo.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}
