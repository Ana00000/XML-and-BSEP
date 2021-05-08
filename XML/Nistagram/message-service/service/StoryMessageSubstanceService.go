package service

import (
	"../model"
	"../repository"
)

type StoryMessageSubstanceService struct {
	Repo *repository.StoryMessageSubstanceRepository
}

func (service * StoryMessageSubstanceService) CreateStoryMessageSubstance(storyMessageSubstance *model.StoryMessageSubstance) error {
	err := service.Repo.CreateStoryMessageSubstance(storyMessageSubstance)
	if err != nil {
		return err
	}
	return nil
}

