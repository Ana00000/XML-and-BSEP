package service

import (
	"../model"
	"../repository"
)

type StoryICRService struct {
	Repo * repository.StoryICRRepository
}

func (service * StoryICRService) CreateStoryICR(storyICR *model.StoryICR) error {
	service.Repo.CreateStoryICR(storyICR)
	return nil
}
