package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
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
