package dto

import "github.com/google/uuid"

type SingleStoryStoryHighlightsDTO struct {
	SingleStoryId uuid.UUID `json:"single_story_id"`
	StoryHighlightId uuid.UUID `json:"story_highlight_id"`
}
