package service

import (
	"../model"
	"../repository"
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
