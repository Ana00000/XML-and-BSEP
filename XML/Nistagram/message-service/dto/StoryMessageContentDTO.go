package dto

import "github.com/google/uuid"

type StoryMessageSubstanceDTO struct {
	Text string `json:"text"`
	StoryId uuid.UUID `json:"story_id"`
}

