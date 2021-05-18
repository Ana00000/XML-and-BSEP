package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type StoryAlbumTagStoryAlbumsRepository struct {
	Database * gorm.DB
}

func (repo * StoryAlbumTagStoryAlbumsRepository) CreateStoryAlbumTagStoryAlbums(storyAlbumTagStoryAlbums *model.StoryAlbumTagStoryAlbums) error {
	result := repo.Database.Create(storyAlbumTagStoryAlbums)
	fmt.Print(result)
	return nil
}