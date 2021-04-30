package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type PostCollectionRepository struct {
	Database * gorm.DB
}

func (repo * PostCollectionRepository) CreatePostCollection(postCollection *model.PostCollection) error {
	result := repo.Database.Create(postCollection)
	fmt.Print(result)
	return nil
}
