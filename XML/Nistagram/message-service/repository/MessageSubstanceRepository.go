package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"gorm.io/gorm"
)

type MessageSubstanceRepository struct {
	Database * gorm.DB
}

func (repo * MessageSubstanceRepository) CreateMessageSubstance(messageSubstance *model.MessageSubstance) error {
	result := repo.Database.Create(messageSubstance)
	fmt.Print(result)
	return nil
}