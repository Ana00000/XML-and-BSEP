package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"fmt"
	"gorm.io/gorm"
)

type ContentRepository struct {
	Database * gorm.DB
}

func (repo * ContentRepository) CreateContent(content *model.Content) error {
	result := repo.Database.Create(content)
	fmt.Print(result)
	return nil
}