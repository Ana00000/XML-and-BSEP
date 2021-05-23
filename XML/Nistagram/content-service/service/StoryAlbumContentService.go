package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
	storyModel "github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
)

type StoryAlbumContentService struct {
	Repo * repository.StoryAlbumContentRepository
}

func (service * StoryAlbumContentService) CreateStoryAlbumContent(storyAlbumContent *model.StoryAlbumContent) error {
	err := service.Repo.CreateStoryAlbumContent(storyAlbumContent)
	if err != nil {
		return err
	}
	return nil
}

func (service *StoryAlbumContentService) FindAllContentsForStoryAlbums(allStoryAlbums []storyModel.StoryAlbum) []model.StoryAlbumContent {
	storyAlbums := service.Repo.FindAllContentsForStoryAlbums(allStoryAlbums)
	if storyAlbums != nil {
		return storyAlbums
	}
	return nil
}

func (service *StoryAlbumContentService) FindAllContentsForStoryAlbum(album *storyModel.StoryAlbum) []model.StoryAlbumContent {
	contents := service.Repo.FindAllContentsForStoryAlbum(album)
	if contents != nil {
		return contents
	}
	return nil
}