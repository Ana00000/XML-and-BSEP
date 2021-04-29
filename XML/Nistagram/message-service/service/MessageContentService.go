package service

import (
	"../model"
	"../repository"
)

type MessageContentService struct {
	Repo *repository.MessageContentRepository
}

func (service * MessageContentService) CreateMessageContent(messageContent *model.MessageContent) error {
	service.Repo.CreateMessageContent(messageContent)
	return nil
}
