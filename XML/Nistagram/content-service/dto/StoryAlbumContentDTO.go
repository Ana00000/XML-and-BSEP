package dto

import (
	"github.com/google/uuid"
)

type StoryAlbumContentDTO struct {
	Path string `json:"path"`
	Type string `json:"type"`
	StoryAlbumId uuid.UUID `json:"story_album_id"`
}
