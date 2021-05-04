package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type CommentTagRepository struct {
	Database * gorm.DB
}

func (repo * CommentTagRepository) CreateCommentTag(commentTag *model.CommentTag) error {
	result := repo.Database.Create(commentTag)
	fmt.Print(result)
	return nil
}
