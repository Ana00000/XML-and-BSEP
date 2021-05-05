package service

import (
	"../model"
	"../repository"
)

type StoryMessageSubstanceService struct {
	Repo *repository.StoryMessageSubstanceRepository
}

func (service * StoryMessageSubstanceService) CreateStoryMessageSubstance(storyMessageSubstance *model.StoryMessageSubstance) error {
	service.Repo.CreateStoryMessageSubstance(storyMessageSubstance)
	return nil
}

