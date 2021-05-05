package dto

import (
	"../model"
	"github.com/google/uuid"
)

type SingleStoryContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	SingleStoryId uuid.UUID `json:"single_story_id"`
}
