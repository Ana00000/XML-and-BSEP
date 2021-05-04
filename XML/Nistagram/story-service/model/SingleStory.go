package model

type SingleStory struct {
	Story
	//StoryContent contentPath.SingleStoryContent `json:"story_content" gorm:"foreignKey:SingleStoryId"`
}
