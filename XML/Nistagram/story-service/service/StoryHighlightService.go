package service

import (
	"../model"
	"../repository"
)

type StoryHighlightService struct {
	Repo * repository.StoryHighlightRepository
}

func (service * StoryHighlightService) CreateStoryHighlight(storyHighlight *model.StoryHighlight) error {
	err := service.Repo.CreateStoryHighlight(storyHighlight)
	if err != nil {
		return err
	}
	return nil
}
