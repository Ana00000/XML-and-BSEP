package dto

import "github.com/google/uuid"

type StoryAlbumContentFullDTO struct {
	ID uuid.UUID `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
	StoryAlbumId uuid.UUID `json:"story_album_id"`
}
