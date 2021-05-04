package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type StoryMessageSubstanceRepository struct {
	Database * gorm.DB
}

func (repo * StoryMessageSubstanceRepository) CreateStoryMessageSubstance(storyMessageSubstance *model.StoryMessageSubstance) error {
	result := repo.Database.Create(storyMessageSubstance)
	fmt.Print(result)
	return nil
}
