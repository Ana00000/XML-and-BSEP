package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type StoryTagStoriesService struct {
	Repo * repository.StoryTagStoriesRepository
}

func (service * StoryTagStoriesService) CreateStoryTagStories(storyTagStories *model.StoryTagStories) error {
	err := service.Repo.CreateStoryTagStories(storyTagStories)
	if err != nil {
		return err
	}
	return nil
}
