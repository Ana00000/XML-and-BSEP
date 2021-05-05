package repository

import (
	"../model"
	"fmt"
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