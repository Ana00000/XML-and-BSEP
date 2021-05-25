package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"fmt"
	"gorm.io/gorm"
)

type StoryRepository struct {
	Database * gorm.DB
}

func (repo * StoryRepository) CreateStory(story *model.Story) error {
	result := repo.Database.Create(story)
	fmt.Print(result)
	return nil
}