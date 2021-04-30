package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type SinglePostRepository struct {
	Database * gorm.DB
}

func (repo * SinglePostRepository) CreateSinglePost(singlePost *model.SinglePost) error {
	result := repo.Database.Create(singlePost)
	fmt.Print(result)
	return nil
}