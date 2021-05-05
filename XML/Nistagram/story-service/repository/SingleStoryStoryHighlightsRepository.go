package repository

import (
	"../model"
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