package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"fmt"
	"gorm.io/gorm"
)

type PostTagPostsRepository struct {
	Database * gorm.DB
}

func (repo * PostTagPostsRepository) CreatePostTagPosts(postTagPosts *model.PostTagPosts) error {
	result := repo.Database.Create(postTagPosts)
	fmt.Print(result)
	return nil
}
