package service

import (
	"../model"
	"../repository"
)

type StoryTagService struct {
	Repo * repository.StoryTagRepository
}

func (service * StoryTagService) CreateStoryTag(storyTag *model.StoryTag) error {
	service.Repo.CreateStoryTag(storyTag)
	return nil
}
