package service

import (
	"../model"
	"../repository"
)

type SinglePostContentService struct {
	Repo * repository.SinglePostContentRepository
}

func (service * SinglePostContentService) CreateSinglePostContent(singlePostContent *model.SinglePostContent) error {
	err := service.Repo.CreateSinglePostContent(singlePostContent)
	if err != nil {
		return err
	}
	return nil
}
