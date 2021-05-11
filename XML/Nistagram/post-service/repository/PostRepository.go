package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
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

func (repo * PostRepository) UpdatePost(post *model.Post) error {
	result := repo.Database.Updates(post)
	fmt.Print(result)
	return nil
}
