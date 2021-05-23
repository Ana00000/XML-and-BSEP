package repository

import (
	"fmt"
	storyModel "github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
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

func (repo *StoryAlbumTagStoryAlbumsRepository) FindAll() []model.StoryAlbumTagStoryAlbums {
	var tags []model.StoryAlbumTagStoryAlbums
	repo.Database.Select("*").Find(&tags)
	return tags
}

func (repo *StoryAlbumTagStoryAlbumsRepository) FindAllTagsForStoryAlbumTagStoryAlbums(albums []storyModel.StoryAlbum) []model.StoryAlbumTagStoryAlbums {
	var tags []model.StoryAlbumTagStoryAlbums
	var allTags = repo.FindAll()

	for i:=0;i<len(albums);i++{
		for j:=0; j<len(allTags);j++{
			if albums[i].ID == allTags[j].StoryAlbumId{
				tags = append(tags, allTags[j])
			}
		}

	}
	return tags
}