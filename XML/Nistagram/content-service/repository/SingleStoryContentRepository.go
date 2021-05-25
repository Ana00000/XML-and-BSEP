package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/dto"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"gorm.io/gorm"
)

type SingleStoryContentRepository struct {
	Database * gorm.DB
}

func (repo * SingleStoryContentRepository) CreateSingleStoryContent(singleStoryContent *model.SingleStoryContent) error {
	result := repo.Database.Create(singleStoryContent)
	fmt.Print(result)
	return nil
}

func (repo *SingleStoryContentRepository) FindAll() []model.SingleStoryContent {
	var stories []model.SingleStoryContent
	repo.Database.Select("*").Find(&stories)
	return stories
}

func (repo *SingleStoryContentRepository) FindAllContentsForStories(allStories []dto.SingleStoryDTO) []model.SingleStoryContent {
	var contents []model.SingleStoryContent
	var allContents = repo.FindAll()

	for i:=0;i<len(allStories);i++{
		for j:=0; j<len(allContents);j++{
			if allStories[i].ID == allContents[j].SingleStoryId{
				contents = append(contents, allContents[j])
			}
		}

	}

	return contents
}

func (repo *SingleStoryContentRepository) FindAllContentsForStory(story *dto.SingleStoryDTO) []model.SingleStoryContent {
	var contents []model.SingleStoryContent
	var allContents = repo.FindAll()

	for j:=0; j<len(allContents);j++{
		if story.ID == allContents[j].SingleStoryId{
			contents = append(contents, allContents[j])
		}
	}

	return contents
}

func (repo * SingleStoryContentRepository) FindSingleStoryContentForStoryId(storyId uuid.UUID) model.SingleStoryContent{
	var singleStoryContent model.SingleStoryContent
	repo.Database.Select("*").Where("single_story_id = ?", storyId).Find(&singleStoryContent)
	return singleStoryContent
}