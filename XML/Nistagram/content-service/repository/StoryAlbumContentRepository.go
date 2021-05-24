package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	storyModel "github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
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

func (repo *StoryAlbumContentRepository) FindAll() []model.StoryAlbumContent {
	var contents []model.StoryAlbumContent
	repo.Database.Select("*").Find(&contents)
	return contents
}

func (repo *StoryAlbumContentRepository) FindAllContentsForStoryAlbums(albums []storyModel.StoryAlbum) []model.StoryAlbumContent {
	var contents []model.StoryAlbumContent
	var allContents = repo.FindAll()

	for i:=0;i<len(albums);i++{
		for j:=0; j<len(allContents);j++{
			if albums[i].ID == allContents[j].StoryAlbumId{
				contents = append(contents, allContents[j])
			}
		}

	}

	return contents
}

func (repo *StoryAlbumContentRepository) FindAllContentsForStoryAlbum(album *storyModel.StoryAlbum) []model.StoryAlbumContent {
	var contents []model.StoryAlbumContent
	var allContents = repo.FindAll()

	for j:=0; j<len(allContents);j++{
		if album.ID == allContents[j].StoryAlbumId{
			contents = append(contents, allContents[j])
		}
	}

	return contents
}