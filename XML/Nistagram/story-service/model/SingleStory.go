package model

import (
	contentPath "../../content-service/model"
)

type SingleStory struct {
	Story
	Content contentPath.Content `json:"content" gorm:"not null"`

}
