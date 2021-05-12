package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"fmt"
	"gorm.io/gorm"
)

type CommentRepository struct {
	Database * gorm.DB
}

func (repo * CommentRepository) CreateComment(comment *model.Comment) error {
	result := repo.Database.Create(comment)
	fmt.Print(result)
	return nil
}