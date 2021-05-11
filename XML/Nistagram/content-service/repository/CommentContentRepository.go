package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"fmt"
	"gorm.io/gorm"
)

type CommentContentRepository struct {
	Database * gorm.DB
}

func (repo * CommentContentRepository) CreateCommentContent(commentContent *model.CommentContent) error {
	result := repo.Database.Create(commentContent)
	fmt.Print(result)
	return nil
}
