package dto

import (
	"github.com/google/uuid"
)

type StoryDTO struct {
	CreationDate string `json:"creationDate"`
	Description string `json:"description"`
	UserId uuid.UUID `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	IsExpired bool `json:"isExpired"`
	Type string `json:"storyType"`
}

