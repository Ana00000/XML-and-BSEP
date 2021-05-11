package dto

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
)

type StoryDTO struct {
	CreationDate string `json:"creationDate"`
	UserId uuid.UUID `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	Type model.StoryType `json:"storyType"`
}

