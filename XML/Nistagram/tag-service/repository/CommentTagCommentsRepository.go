package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type CommentTagCommentsRepository struct {
	Database * gorm.DB
}

func (repo * CommentTagCommentsRepository) CreateCommentTagComments(commentTagComments *model.CommentTagComments) error {
	result := repo.Database.Create(commentTagComments)
	fmt.Print(result)
	return nil
}
