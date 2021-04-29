package service

import (
	"../model"
	"../repository"
)

type StoryHighlightService struct {
	Repo * repository.StoryHighlightRepository
}

func (service * StoryHighlightService) CreateStoryHighlight(storyHighlight *model.StoryHighlight) error {
	service.Repo.CreateStoryHighlight(storyHighlight)
	return nil
}
