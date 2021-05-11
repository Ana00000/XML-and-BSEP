package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
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
