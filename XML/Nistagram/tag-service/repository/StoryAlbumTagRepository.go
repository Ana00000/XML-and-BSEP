package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type StoryAlbumTagRepository struct {
	Database * gorm.DB
}

func (repo * StoryAlbumTagRepository) CreateStoryAlbumTag(storyAlbumTag *model.StoryAlbumTag) error {
	result := repo.Database.Create(storyAlbumTag)
	fmt.Print(result)
	return nil
}