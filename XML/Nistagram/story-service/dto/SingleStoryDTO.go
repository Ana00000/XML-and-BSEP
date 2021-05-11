package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/google/uuid"
)

type SingleStoryDTO struct {
	CreationDate string `json:"creationDate"`
	UserId uuid.UUID `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	Type model.StoryType `json:"storyType"`
}