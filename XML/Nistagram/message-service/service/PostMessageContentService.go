package service

import (
	"../model"
	"../repository"
)

type PostMessageContentService struct {
	Repo *repository.PostMessageContentRepository
}

func (service * PostMessageContentService) CreatePostMessageContent(postMessageContent *model.PostMessageContent) error {
	service.Repo.CreatePostMessageContent(postMessageContent)
	return nil
}


