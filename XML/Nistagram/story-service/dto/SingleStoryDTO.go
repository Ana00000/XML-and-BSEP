package dto

import (
	"github.com/google/uuid"
)

type SingleStoryDTO struct {
	CreationDate string `json:"creationDate"`
	Description string `json:"description"`
	UserId uuid.UUID `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	Type string `json:"storyType"`
}