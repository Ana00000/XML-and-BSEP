package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/repository"
)

type SingleStoryStoryHighlightsService struct {
	Repo * repository.SingleStoryStoryHighlightsRepository
}

func (service * SingleStoryStoryHighlightsService) CreateSingleStoryStoryHighlights(singleStoryStoryHighlights *model.SingleStoryStoryHighlights) error {
	err := service.Repo.CreateSingleStoryStoryHighlights(singleStoryStoryHighlights)
	if err != nil {
		return err
	}
	return nil
}

func (service * SingleStoryStoryHighlightsService) FindAllSingleStoryStoryHighlightsForStory(storyId uuid.UUID) []model.SingleStoryStoryHighlights{
	singleStoryStoryHighlights := service.Repo.FindAllSingleStoryStoryHighlightsForStory(storyId)
	if singleStoryStoryHighlights != nil {
		return singleStoryStoryHighlights
	}
	return nil
}

func (service * SingleStoryStoryHighlightsService) FindAllSingleStoryStoryHighlightsForStoryHighlight(storyHighlightId uuid.UUID) []model.SingleStoryStoryHighlights{
	singleStoryStoryHighlights := service.Repo.FindAllSingleStoryStoryHighlightsForStoryHighlight(storyHighlightId)
	if singleStoryStoryHighlights != nil {
		return singleStoryStoryHighlights
	}
	return nil
}