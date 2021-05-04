package service

import (
	"../model"
	"../repository"
)

type PostTagService struct {
	Repo * repository.PostTagRepository
}

func (service * PostTagService) CreatePostTag(postTag *model.PostTag) error {
	service.Repo.CreatePostTag(postTag)
	return nil
}
