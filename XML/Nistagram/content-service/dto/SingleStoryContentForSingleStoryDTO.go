package dto

import "github.com/google/uuid"

type SingleStoryContentForSingleStoryDTO struct {
	ID uuid.UUID `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
	SingleStoryId uuid.UUID `json:"single_story_id"`
}