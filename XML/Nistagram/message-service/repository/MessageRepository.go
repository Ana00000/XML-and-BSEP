package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
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
