package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"fmt"
	"gorm.io/gorm"
)

type SingleStoryContentRepository struct {
	Database * gorm.DB
}

func (repo * SingleStoryContentRepository) CreateSingleStoryContent(singleStoryContent *model.SingleStoryContent) error {
	result := repo.Database.Create(singleStoryContent)
	fmt.Print(result)
	return nil
}
