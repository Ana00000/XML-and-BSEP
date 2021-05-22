package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type StoryTagService struct {
	Repo * repository.StoryTagRepository
}

func (service * StoryTagService) CreateStoryTag(storyTag *model.StoryTag) error {
	err := service.Repo.CreateStoryTag(storyTag)
	if err != nil {
		return err
	}
	return nil
}

func (service * StoryTagService) FindStoryTagForId(storyId uuid.UUID) model.StoryTag{
	storyTag := service.Repo.FindStoryTagForId(storyId)
	return storyTag
}