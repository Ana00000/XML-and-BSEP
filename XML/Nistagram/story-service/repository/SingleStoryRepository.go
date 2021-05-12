package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"fmt"
	"gorm.io/gorm"
)

type SingleStoryRepository struct {
	Database * gorm.DB
}

func (repo * SingleStoryRepository) CreateSingleStory(singleStory *model.SingleStory) error {
	result := repo.Database.Create(singleStory)
	fmt.Print(result)
	return nil
}