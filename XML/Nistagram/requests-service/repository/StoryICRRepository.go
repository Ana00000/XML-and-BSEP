package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"fmt"
	"gorm.io/gorm"
)

type StoryICRRepository struct {
	Database * gorm.DB
}

func (repo * StoryICRRepository) CreateStoryICR(storyICR *model.StoryICR) error {
	result := repo.Database.Create(storyICR)
	fmt.Print(result)
	return nil
}
