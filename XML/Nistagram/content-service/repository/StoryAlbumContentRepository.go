package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type StoryAlbumContentRepository struct {
	Database * gorm.DB
}

func (repo * StoryAlbumContentRepository) CreateStoryAlbumContent(storyAlbumContent *model.StoryAlbumContent) error {
	result := repo.Database.Create(storyAlbumContent)
	fmt.Print(result)
	return nil
}
