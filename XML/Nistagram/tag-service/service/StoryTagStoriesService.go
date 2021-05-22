package service

import (
	"github.com/google/uuid"
	storyModel "github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type StoryTagStoriesService struct {
	Repo * repository.StoryTagStoriesRepository
}

func (service * StoryTagStoriesService) CreateStoryTagStories(storyTagStories *model.StoryTagStories) error {
	err := service.Repo.CreateStoryTagStories(storyTagStories)
	if err != nil {
		return err
	}
	return nil
}

func (service *StoryTagStoriesService) FindAllTagsForStories(allStories []storyModel.SingleStory) []model.StoryTagStories {
	tags := service.Repo.FindAllTagsForStories(allStories)
	if tags != nil {
		return tags
	}
	return nil
}

func (service *StoryTagStoriesService) FindAllTagsForStory(story *storyModel.SingleStory) []model.StoryTagStories {
	tags := service.Repo.FindAllTagsForStory(story)
	if tags != nil {
		return tags
	}
	return nil
}

func (service * StoryTagStoriesService) FindStoryTagStoriesForStoryId(storyId uuid.UUID) []model.StoryTagStories{
	storyTagStories := service.Repo.FindStoryTagStoriesForStoryId(storyId)
	if storyTagStories != nil {
		return storyTagStories
	}
	return nil
}