package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"fmt"
	"gorm.io/gorm"
)

type StoryTagRepository struct {
	Database * gorm.DB
}

func (repo * StoryTagRepository) CreateStoryTag(storyTag *model.StoryTag) error {
	result := repo.Database.Create(storyTag)
	fmt.Print(result)
	return nil
}
