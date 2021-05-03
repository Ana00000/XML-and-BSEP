package model

import (
	storyPath "../../story-service/model"
)

type StoryTag struct {
	Tag
	Stories []storyPath.Story `gorm:"many2many:story_tag_stories"`
}
