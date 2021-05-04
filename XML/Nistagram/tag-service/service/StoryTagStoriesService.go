package service

import (
	"../model"
	"../repository"
)

type StoryTagStoriesService struct {
	Repo * repository.StoryTagStoriesRepository
}

func (service * StoryTagStoriesService) CreateStoryTagStories(storyTagStories *model.StoryTagStories) error {
	service.Repo.CreateStoryTagStories(storyTagStories)
	return nil
}
