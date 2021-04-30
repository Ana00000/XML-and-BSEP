package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type PostICRRepository struct {
	Database * gorm.DB
}

func (repo * PostICRRepository) CreatePostICR(postICR *model.PostICR) error {
	result := repo.Database.Create(postICR)
	fmt.Print(result)
	return nil
}
