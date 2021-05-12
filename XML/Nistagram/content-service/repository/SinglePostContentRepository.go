package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"fmt"
	"gorm.io/gorm"
)

type SinglePostContentRepository struct {
	Database * gorm.DB
}

func (repo * SinglePostContentRepository) CreateSinglePostContent(singlePostContent *model.SinglePostContent) error {
	result := repo.Database.Create(singlePostContent)
	fmt.Print(result)
	return nil
}
