package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type TagRepository struct {
	Database * gorm.DB
}

func (repo * TagRepository) CreateTag(tag *model.Tag) error {
	result := repo.Database.Create(tag)
	fmt.Print(result)
	return nil
}
