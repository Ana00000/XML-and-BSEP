package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"fmt"
	"gorm.io/gorm"
)

type StoryTagStoriesRepository struct {
	Database * gorm.DB
}

func (repo * StoryTagStoriesRepository) CreateStoryTagStories(storyTagStories *model.StoryTagStories) error {
	result := repo.Database.Create(storyTagStories)
	fmt.Print(result)
	return nil
}