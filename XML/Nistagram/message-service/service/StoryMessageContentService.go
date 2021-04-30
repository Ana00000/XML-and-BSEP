package service

import (
	"../model"
	"../repository"
)

type StoryMessageContentService struct {
	Repo *repository.StoryMessageContentRepository
}

func (service * StoryMessageContentService) CreateStoryMessageContent(storyMessageContent *model.StoryMessageContent) error {
	service.Repo.CreateStoryMessageContent(storyMessageContent)
	return nil
}

