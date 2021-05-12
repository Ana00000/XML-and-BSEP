package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/repository"
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

