package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type MessageContentRepository struct {
	Database * gorm.DB
}

func (repo * MessageContentRepository) CreateMessageContent(messageContent *model.MessageContent) error {
	result := repo.Database.Create(messageContent)
	fmt.Print(result)
	return nil
}