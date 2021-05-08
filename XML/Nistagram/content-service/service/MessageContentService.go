package service

import (
	"../model"
	"../repository"
)

type MessageContentService struct {
	Repo * repository.MessageContentRepository
}

func (service * MessageContentService) CreateMessageContent(messageContent *model.MessageContent) error {
	err := service.Repo.CreateMessageContent(messageContent)
	if err != nil {
		return err
	}
	return nil
}
