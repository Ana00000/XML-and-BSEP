package repository

import (
	"fmt"
	"github.com/google/uuid"
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

func (repo * PostCollectionPostsRepository) FindAllPostCollectionPostsForPost(postId uuid.UUID) []model.PostCollectionPosts{
	var postCollectionPosts []model.PostCollectionPosts
	repo.Database.Select("*").Where("single_post_id = ?", postId).Find(&postCollectionPosts)
	return postCollectionPosts
}