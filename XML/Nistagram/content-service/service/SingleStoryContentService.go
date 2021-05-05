package service

import (
	"../model"
	"../repository"
)

type SingleStoryContentService struct {
	Repo * repository.SingleStoryContentRepository
}

func (service * SingleStoryContentService) CreateSingleStoryContent(singleStoryContent *model.SingleStoryContent) error {
	service.Repo.CreateSingleStoryContent(singleStoryContent)
	return nil
}
