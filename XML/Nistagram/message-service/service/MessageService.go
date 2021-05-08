package service

import (
	"../model"
	"../repository"
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