package dto

import "github.com/google/uuid"

type StoryTagStoriesDTO struct {
	StoryTagId uuid.UUID `json:"story_tag_id"`
	StoryId uuid.UUID `json:"story_id"`
}

