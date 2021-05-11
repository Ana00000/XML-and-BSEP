package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
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