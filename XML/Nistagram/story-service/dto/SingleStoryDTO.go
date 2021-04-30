package dto

import (
	contentPath "../../content-service/model"
	"../model"
	"github.com/google/uuid"
)

type SingleStoryDTO struct {
	CreationDate string `json:"creationDate"`
	UserId string `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	Type model.StoryType `json:"storyType"`
	Content contentPath.Content `json:"content"`
}