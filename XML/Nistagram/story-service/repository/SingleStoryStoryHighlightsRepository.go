package repository

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"fmt"
	"gorm.io/gorm"
)

type SingleStoryStoryHighlightsRepository struct {
	Database * gorm.DB
}

func (repo * SingleStoryStoryHighlightsRepository) CreateSingleStoryStoryHighlights(singleStoryStoryHighlights *model.SingleStoryStoryHighlights) error {
	result := repo.Database.Create(singleStoryStoryHighlights)
	fmt.Print(result)
	return nil
}

func (repo * SingleStoryStoryHighlightsRepository) FindAllSingleStoryStoryHighlightsForStory(storyId uuid.UUID) []model.SingleStoryStoryHighlights{
	var singleStoryStoryHighlights []model.SingleStoryStoryHighlights
	repo.Database.Select("*").Where("single_story_id = ?", storyId).Find(&singleStoryStoryHighlights)
	return singleStoryStoryHighlights
}