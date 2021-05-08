package service

import (
	"../model"
	"../repository"
)

type PostTagService struct {
	Repo * repository.PostTagRepository
}

func (service * PostTagService) CreatePostTag(postTag *model.PostTag) error {
	err := service.Repo.CreatePostTag(postTag)
	if err != nil {
		return err
	}
	return nil
}
