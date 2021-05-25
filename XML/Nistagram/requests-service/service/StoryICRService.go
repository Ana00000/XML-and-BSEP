package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
)

type StoryICRService struct {
	Repo * repository.StoryICRRepository
}

func (service * StoryICRService) CreateStoryICR(storyICR *model.StoryICR) error {
	err := service.Repo.CreateStoryICR(storyICR)
	if err != nil {
		return err
	}
	return nil
}
