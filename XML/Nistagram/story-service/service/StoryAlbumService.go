package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/repository"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
)

type StoryAlbumService struct {
	Repo * repository.StoryAlbumRepository
}

func (service * StoryAlbumService) CreateStoryAlbum(storyAlbum *model.StoryAlbum) error {
	err := service.Repo.CreateStoryAlbum(storyAlbum)
	if err != nil {
		return err
	}
	return nil
}

func (service *StoryAlbumService) FindAllAlbumStoriesForUser(ID uuid.UUID) []model.StoryAlbum {
	albumStories := service.Repo.FindAllAlbumStoriesForUser(ID)
	if albumStories != nil {
		return albumStories
	}
	return nil
}

func (service *StoryAlbumService) FindByID(ID uuid.UUID) *model.StoryAlbum {
	storyAlbum := service.Repo.FindByID(ID)
	return storyAlbum
}

func (service *StoryAlbumService) FindAllPublicAlbumStoriesNotRegisteredUser(users []userModel.ClassicUser) []model.StoryAlbum {
	storyAlbums := service.Repo.FindAllPublicAlbumStoriesNotRegisteredUser(users)
	if storyAlbums != nil {
		return storyAlbums
	}
	return nil
}

// FIND ALL NOT DELETED VALID STORY ALBUMS THAT LOGGED IN USER FOLLOWS
func (service *StoryAlbumService) FindAllFollowingStoryAlbums(followings []userModel.ClassicUserFollowings) []model.StoryAlbum {
	storyAlbums := service.Repo.FindAllFollowingStoryAlbums(followings)
	if storyAlbums != nil {
		return storyAlbums
	}
	return nil
}