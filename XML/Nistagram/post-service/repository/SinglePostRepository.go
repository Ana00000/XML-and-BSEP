package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
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