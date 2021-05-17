package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type PostTagRepository struct {
	Database * gorm.DB
}

func (repo * PostTagRepository) CreatePostTag(postTag *model.PostTag) error {
	result := repo.Database.Create(postTag)
	fmt.Print(result)
	return nil
}
