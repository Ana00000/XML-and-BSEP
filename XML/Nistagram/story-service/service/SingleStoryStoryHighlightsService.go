package service

import (
	"../model"
	"../repository"
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
