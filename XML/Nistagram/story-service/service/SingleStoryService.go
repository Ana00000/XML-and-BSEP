package service

import (
	"../model"
	"../repository"
)

type SingleStoryService struct {
	Repo * repository.SingleStoryRepository
}

func (service * SingleStoryService) CreateSingleStory(singleStory *model.SingleStory) error {
	err := service.Repo.CreateSingleStory(singleStory)
	if err != nil {
		return err
	}
	return nil
}