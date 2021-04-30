package dto

import (
	"../model"
	"github.com/google/uuid"
)


type StoryDTO struct {
	CreationDate string `json:"creationDate"`
	UserId string `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	Type model.StoryType `json:"storyType"`
}

