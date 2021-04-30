package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type StoryMessageContentRepository struct {
	Database * gorm.DB
}

func (repo * StoryMessageContentRepository) CreateStoryMessageContent(storyMessageContent *model.StoryMessageContent) error {
	result := repo.Database.Create(storyMessageContent)
	fmt.Print(result)
	return nil
}
