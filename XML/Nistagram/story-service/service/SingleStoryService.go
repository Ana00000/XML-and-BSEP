package service

import (
	"../model"
	"../repository"
)

type SingleStoryService struct {
	Repo * repository.SingleStoryRepository
}

func (service * SingleStoryService) CreateSingleStory(singleStory *model.SingleStory) error {
	service.Repo.CreateSingleStory(singleStory)
	return nil
}