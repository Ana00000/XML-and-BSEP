package dto

import "github.com/google/uuid"

type SingleStoryContentDTO struct {
	ID uuid.UUID `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
	SingleStoryId uuid.UUID `json:"single_story_id"`
}
