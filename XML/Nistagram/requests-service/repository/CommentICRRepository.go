package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"fmt"
	"gorm.io/gorm"
)

type CommentICRRepository struct {
	Database * gorm.DB
}

func (repo * CommentICRRepository) CreateCommentICR(commentICR *model.CommentICR) error {
	result := repo.Database.Create(commentICR)
	fmt.Print(result)
	return nil
}
