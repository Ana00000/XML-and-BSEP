package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
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
