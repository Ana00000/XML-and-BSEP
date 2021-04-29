package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type MessageRepository struct {
	Database * gorm.DB
}

func (repo * MessageRepository) CreateMessage(message *model.Message) error {
	result := repo.Database.Create(message)
	fmt.Print(result)
	return nil
}
