package dto

import (
	"github.com/google/uuid"
)

type SingleStoryContentDTO struct {
	Path string `json:"path"`
	Type string `json:"type"`
	SingleStoryId uuid.UUID `json:"single_story_id"`
}
