package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/repository"
)

type MessageService struct {
	Repo * repository.MessageRepository
}

func (service * MessageService) CreateMessage(message *model.Message) error {
	err := service.Repo.CreateMessage(message)
	if err != nil {
		return err
	}
	return nil
}