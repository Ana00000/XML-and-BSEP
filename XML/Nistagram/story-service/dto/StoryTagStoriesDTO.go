package dto

import (
	"github.com/google/uuid"
)

type StoryTagStoriesDTO struct {
	ID uuid.UUID `json:"id"`
	TagId uuid.UUID `json:"tag_id"`
	StoryId uuid.UUID `json:"story_id"`
}
