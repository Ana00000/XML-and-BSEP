package dto

import (
	locationPath "../../location-service/model"
	"../model"
	"time"
)


type StoryDTO struct {
	CreationDate time.Time `json: "creationDate"`
	UserId string `json:"userId"`
	Location locationPath.Location `json:"location"`
	IsDeleted bool `json:"isDeleted"`
	Type model.StoryType `json:"storyType"`
}

