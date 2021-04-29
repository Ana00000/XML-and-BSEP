package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type StoryHighlightRepository struct {
	Database * gorm.DB
}

func (repo * StoryHighlightRepository) CreateStoryHighlight(storyHighlight *model.StoryHighlight) error {
	result := repo.Database.Create(storyHighlight)
	fmt.Print(result)
	return nil
}
