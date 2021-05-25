package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
)

type SingleStoryContentService struct {
	Repo * repository.SingleStoryContentRepository
}

func (service * SingleStoryContentService) CreateSingleStoryContent(singleStoryContent *model.SingleStoryContent) error {
	err := service.Repo.CreateSingleStoryContent(singleStoryContent)
	if err != nil {
		return err
	}
	return nil
}

func (service *SingleStoryContentService) FindAllContentsForStories(allStories []dto.SingleStoryDTO) []model.SingleStoryContent {
	stories := service.Repo.FindAllContentsForStories(allStories)
	if stories != nil {
		return stories
	}
	return nil
}

func (service *SingleStoryContentService) FindAllContentsForStory(story *dto.SingleStoryDTO) []model.SingleStoryContent {
	stories := service.Repo.FindAllContentsForStory(story)
	if stories != nil {
		return stories
	}
	return nil
}

func (service * SingleStoryContentService) FindSingleStoryContentForStoryId(storyId uuid.UUID) model.SingleStoryContent{
	singleStoryContent := service.Repo.FindSingleStoryContentForStoryId(storyId)
	return singleStoryContent
}