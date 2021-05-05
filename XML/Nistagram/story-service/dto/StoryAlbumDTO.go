package dto

import (
	"../model"
	"github.com/google/uuid"
)

type StoryAlbumDTO struct {
	CreationDate string `json:"creationDate"`
	UserId uuid.UUID `json:"userId"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted"`
	Type model.StoryType `json:"storyType"`
}
