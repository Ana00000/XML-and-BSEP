package repository

import (
	"fmt"
	"github.com/google/uuid"
	storyModel "github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"gorm.io/gorm"
)

type StoryTagStoriesRepository struct {
	Database * gorm.DB
}

func (repo * StoryTagStoriesRepository) CreateStoryTagStories(storyTagStories *model.StoryTagStories) error {
	result := repo.Database.Create(storyTagStories)
	fmt.Print(result)
	return nil
}

func (repo *StoryTagStoriesRepository) FindAll() []model.StoryTagStories {
	var tags []model.StoryTagStories
	repo.Database.Select("*").Find(&tags)
	return tags
}


func (repo *StoryTagStoriesRepository) FindAllTagsForStories(allStories []storyModel.SingleStory) []model.StoryTagStories {
	var tags []model.StoryTagStories
	var allTags = repo.FindAll()

	for i:=0;i<len(allStories);i++{
		for j:=0; j<len(allTags);j++{
			if allStories[i].ID == allTags[j].StoryId{
				tags = append(tags, allTags[j])
			}
		}

	}
	return tags
}

func (repo *StoryTagStoriesRepository) FindAllTagsForStory(story *storyModel.SingleStory) []model.StoryTagStories {
	var tags []model.StoryTagStories
	var allTags = repo.FindAll()

	for j:=0; j<len(allTags);j++{
		if story.ID == allTags[j].StoryId{
			tags = append(tags, allTags[j])
		}
	}

	return tags
}

func (repo * StoryTagStoriesRepository) FindStoryTagStoriesForStoryId(storyId uuid.UUID) []model.StoryTagStories{
	var storyTagStories []model.StoryTagStories
	repo.Database.Select("*").Where("story_id = ?", storyId).Find(&storyTagStories)
	return storyTagStories
}