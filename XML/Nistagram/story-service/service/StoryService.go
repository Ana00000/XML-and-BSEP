package service

import (
	"../model"
	"../repository"
)

type StoryService struct {
	Repo * repository.StoryRepository
}

func (service * StoryService) CreateStory(story *model.Story) error {
	service.Repo.CreateStory(story)
	return nil
}