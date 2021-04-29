package dto

import (
	contentPath "../../content-service/model"
	locationPath "../../location-service/model"
	"../model"
	"time"
)

type SingleStoryDTO struct {
	CreationDate time.Time `json: "creationDate"`
	UserId string `json:"userId"`
	Location locationPath.Location `json:"location"`
	IsDeleted bool `json:"isDeleted"`
	Type model.StoryType `json:"storyType"`
	Content contentPath.Content `json:"content"`
}