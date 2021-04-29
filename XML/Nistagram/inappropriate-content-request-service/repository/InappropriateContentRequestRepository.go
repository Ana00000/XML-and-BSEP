package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type InappropriateContentRequestRepository struct {
	Database * gorm.DB
}

func (repo * InappropriateContentRequestRepository) CreateInappropriateContentRequest(inappropriateContentRequest *model.InappropriateContentRequest) error {
	result := repo.Database.Create(inappropriateContentRequest)
	fmt.Print(result)
	return nil
}
