package dto

import (
	"github.com/google/uuid"
)

type StoryDTO struct {
	CreationDate string `json:"creationDate"`
	UserId uuid.UUID `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	Type string `json:"storyType"`
}

