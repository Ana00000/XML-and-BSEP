package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/repository"
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

func (service * StoryHighlightService) FindAllStoryHighlightsForUser(userId uuid.UUID) []model.StoryHighlight{
	storyHighlights := service.Repo.FindAllStoryHighlightsForUser(userId)
	if storyHighlights != nil {
		return storyHighlights
	}
	return nil
}