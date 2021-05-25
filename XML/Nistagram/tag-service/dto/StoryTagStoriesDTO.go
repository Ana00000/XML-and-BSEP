package dto

import "github.com/google/uuid"

type StoryTagStoriesDTO struct {
	TagId uuid.UUID `json:"tag_id"`
	StoryId uuid.UUID `json:"story_id"`
}

