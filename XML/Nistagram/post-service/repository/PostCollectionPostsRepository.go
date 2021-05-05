package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type PostCollectionPostsRepository struct {
	Database * gorm.DB
}

func (repo * PostCollectionPostsRepository) CreatePostCollectionPosts(postCollectionPosts *model.PostCollectionPosts) error {
	result := repo.Database.Create(postCollectionPosts)
	fmt.Println(result)
	return nil
}