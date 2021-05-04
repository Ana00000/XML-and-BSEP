package model

type StoryAlbum struct {
	Story
	//StoryContents contentPath.SingleStoryContent `json:"story_content" gorm:"foreignKey:SingleStoryId"`
}
