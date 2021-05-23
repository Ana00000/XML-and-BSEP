package repository

import (
	"github.com/google/uuid"
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

func (repo *StoryAlbumRepository) FindAllAlbumStoriesForUser(userId uuid.UUID) []model.StoryAlbum {
	var storyAlbums []model.StoryAlbum
	repo.Database.Select("*").Where("user_id = ? and is_deleted = ?", userId, false).Find(&storyAlbums)
	return storyAlbums
}