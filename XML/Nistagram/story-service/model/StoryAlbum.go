package model

import (
	contentPath "../../content-service/model"
)

type StoryAlbum struct {
	Story
	StoryContents contentPath.SingleStoryContent `json:"story_content" gorm:"foreignKey:SingleStoryId"`
}
