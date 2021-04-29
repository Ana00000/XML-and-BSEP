package service

import (
	"../model"
	"../repository"
)

type MessageService struct {
	Repo * repository.MessageRepository
}

func (service * MessageService) CreateMessage(message *model.Message) error {
	service.Repo.CreateMessage(message)
	return nil
}