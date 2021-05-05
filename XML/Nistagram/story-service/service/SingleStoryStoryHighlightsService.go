package service

import (
	"../model"
	"../repository"
)

type SingleStoryStoryHighlightsService struct {
	Repo * repository.SingleStoryStoryHighlightsRepository
}

func (service * SingleStoryStoryHighlightsService) CreateSingleStoryStoryHighlights(singleStoryStoryHighlights *model.SingleStoryStoryHighlights) error {
	service.Repo.CreateSingleStoryStoryHighlights(singleStoryStoryHighlights)
	return nil
}
