package dto

import "github.com/google/uuid"

type StoryICRDTO struct {
	Note string `json:"note"`
	UserId uuid.UUID `json:"userId"`
	StoryId uuid.UUID `json:"storyId"`
}