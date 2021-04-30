package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type PostRepository struct {
	Database * gorm.DB
}

func (repo * PostRepository) CreatePost(post *model.Post) error {
	result := repo.Database.Create(post)
	fmt.Print(result)
	return nil
}
