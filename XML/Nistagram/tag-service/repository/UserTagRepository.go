package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type UserTagRepository struct {
	Database * gorm.DB
}

func (repo * UserTagRepository) CreateUserTag(userTag *model.UserTag) error {
	result := repo.Database.Create(userTag)
	fmt.Print(result)
	return nil
}
