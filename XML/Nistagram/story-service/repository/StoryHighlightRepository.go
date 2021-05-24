package repository

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
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

func (repo * StoryHighlightRepository) FindAllStoryHighlightsForUser(userId uuid.UUID) []model.StoryHighlight{
	var storyHighlights []model.StoryHighlight
	repo.Database.Select("*").Where("user_id = ?", userId).Find(&storyHighlights)
	return storyHighlights
}