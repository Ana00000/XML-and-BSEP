package model

type StoryICR struct {
	InappropriateContentRequest
	StoryId string `json: "storyId" gorm:"not null"`
}