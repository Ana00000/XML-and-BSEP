package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"fmt"
	"gorm.io/gorm"
)

type StoryAlbumRepository struct {
	Database * gorm.DB
}

func (repo * StoryAlbumRepository) CreateStoryAlbum(storyAlbum *model.StoryAlbum) error {
	result := repo.Database.Create(storyAlbum)
	fmt.Print(result)
	return nil
}