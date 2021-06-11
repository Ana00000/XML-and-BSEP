package dto

import "github.com/google/uuid"

type StoryICRDTO struct {
	Note    string    `json:"note" validate:"required"`
	UserId  uuid.UUID `json:"userId" validate:"required"`
	StoryId uuid.UUID `json:"storyId" validate:"required"`
}
