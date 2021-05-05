package dto

import (
	"../model"
	"github.com/google/uuid"
)

type StoryAlbumContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	StoryAlbumId uuid.UUID `json:"story_album_id"`
}
