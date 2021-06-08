package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
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

func (repo CommentTagCommentsRepository) FindAllCommentTagCommentsByCommentId(id uuid.UUID) []model.CommentTagComments {
	var tags []model.CommentTagComments
	repo.Database.Select("*").Where("comment_id = ?", id).Find(&tags)
	return tags
}
