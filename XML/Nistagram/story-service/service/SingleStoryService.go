package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/dto"
)

type SingleStoryService struct {
	Repo * repository.SingleStoryRepository
}

func (service * SingleStoryService) CreateSingleStory(singleStory *model.SingleStory) error {
	err := service.Repo.CreateSingleStory(singleStory)
	if err != nil {
		return err
	}
	return nil
}


func (service *SingleStoryService) FindByID(ID uuid.UUID) *model.SingleStory {
	story := service.Repo.FindByID(ID)
	return story
}

// NOT REGISTERED
func (service *SingleStoryService) FindAllStoriesForUserNotReg(ID uuid.UUID) []model.SingleStory {
	stories := service.Repo.FindAllStoriesForUserNotReg(ID)
	if stories != nil {
		return stories
	}
	return nil
}

func (service *SingleStoryService) FindAllPublicStoriesNotRegisteredUser(allValidUsers []dto.ClassicUserDTO) []model.SingleStory {
	stories := service.Repo.FindAllPublicStoriesNotRegisteredUser(allValidUsers)
	if stories != nil {
		return stories
	}
	return nil
}

// REGISTERED


// metoda koja se poziva kada registrovani user udje na profil nekog usera KOME SE NALAZI U CLOSE FRIENDS
// ZNACI USLOVI: CLOSE FIRENDS ILI ALL FRIENDS ILI PUBLIC STORY
func (service *SingleStoryService) FindAllStoriesForUserCloseFriend(ID uuid.UUID) []model.SingleStory {
	stories := service.Repo.FindAllStoriesForUserCloseFriend(ID)
	if stories != nil {
		return stories
	}
	return nil
}


// metoda koja se poziva kada registrovani user udje na profil nekog usera koga prati (PUBLIC ILI PRIVATE) ali nije CLOSE FRIENDS
// ZNACI USLOVI: ILI PUBLIC STORY ILI ALL FRIENDS STORY
func (service *SingleStoryService) FindAllStoriesForUserPublicAllFriends(ID uuid.UUID) []model.SingleStory {
	stories := service.Repo.FindAllStoriesForUserPublicAllFriends(ID)
	if stories != nil {
		return stories
	}
	return nil
}

// metoda koja se poziva kada registrovani user udje na profil nekog usera koga ne prati ali je user PUBLIC
// ZNACI USLOVI: ILI PUBLIC STORY
func (service *SingleStoryService) FindAllStoriesForUserPublic(ID uuid.UUID) []model.SingleStory {
	stories := service.Repo.FindAllStoriesForUserPublic(ID)
	if stories != nil {
		return stories
	}
	return nil
}

// FIND ALL NOT DELETED VALID STORIES THAT LOGGED IN USER FOLLOWS
func (service *SingleStoryService) FindAllFollowingStories(followings []dto.ClassicUserFollowingsDTO) []model.SingleStory {
	stories := service.Repo.FindAllFollowingStories(followings)
	if stories != nil {
		return stories
	}
	return nil
}

// FIND ALL NOT DELETED STORIES FROM MY LOGGED IN USER (WITH EXPIRED STORIES)
func (service *SingleStoryService) FindAllStoriesForLoggedUser(userId uuid.UUID) []model.SingleStory {
	stories := service.Repo.FindAllStoriesForLoggedUser(userId)
	if stories != nil {
		return stories
	}
	return nil
}

func (service * SingleStoryService) FindSingleStoryForId(storyId uuid.UUID) model.SingleStory{
	singleStory := service.Repo.FindSingleStoryForId(storyId)
	return singleStory
}