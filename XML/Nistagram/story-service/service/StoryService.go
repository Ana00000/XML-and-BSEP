package service

import (
	"../model"
	"../repository"
)

type StoryService struct {
	Repo * repository.StoryRepository
}

func (service * StoryService) CreateStory(story *model.Story) error {
	err := service.Repo.CreateStory(story)
	if err != nil {
		return err
	}
	return nil
}