package service

import (
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
