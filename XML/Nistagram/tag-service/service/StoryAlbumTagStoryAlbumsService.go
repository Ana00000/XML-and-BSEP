package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
)

type StoryAlbumTagStoryAlbumsService struct {
	Repo * repository.StoryAlbumTagStoryAlbumsRepository
}

func (service * StoryAlbumTagStoryAlbumsService) CreateStoryAlbumTagStoryAlbums(storyAlbumTagStoryAlbums *model.StoryAlbumTagStoryAlbums) error {
	err := service.Repo.CreateStoryAlbumTagStoryAlbums(storyAlbumTagStoryAlbums)
	if err != nil {
		return err
	}
	return nil
}

func (service *StoryAlbumTagStoryAlbumsService) FindAllTagsForStoryAlbumTagStoryAlbums(albums []dto.StoryAlbumFullDTO) []model.StoryAlbumTagStoryAlbums {
	tags := service.Repo.FindAllTagsForStoryAlbumTagStoryAlbums(albums)
	if tags != nil {
		return tags
	}
	return nil
}

func (service *StoryAlbumTagStoryAlbumsService) FindAllTagsForStoryAlbum(album *dto.StoryAlbumFullDTO) []model.StoryAlbumTagStoryAlbums {
	tags := service.Repo.FindAllTagsForStoryAlbum(album)
	if tags != nil {
		return tags
	}
	return nil
}