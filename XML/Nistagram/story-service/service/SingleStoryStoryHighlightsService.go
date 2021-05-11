package service

import (
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
