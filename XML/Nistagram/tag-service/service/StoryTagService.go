package service

import (
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
