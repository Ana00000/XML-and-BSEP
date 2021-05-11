package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/repository"
)

type StoryService struct {
	Repo * repository.StoryRepository
}

func (service * StoryService) CreateStory(story *model.Story) error {
	err := service.Repo.CreateStory(story)
	if err != nil {
		return err
	}
	return nil
}