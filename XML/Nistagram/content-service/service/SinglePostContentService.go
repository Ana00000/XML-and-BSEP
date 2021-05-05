package service

import (
	"../model"
	"../repository"
)

type SinglePostContentService struct {
	Repo * repository.SinglePostContentRepository
}

func (service * SinglePostContentService) CreateSinglePostContent(singlePostContent *model.SinglePostContent) error {
	service.Repo.CreateSinglePostContent(singlePostContent)
	return nil
}
