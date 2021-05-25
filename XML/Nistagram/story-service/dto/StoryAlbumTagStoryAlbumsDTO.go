package dto

import "github.com/google/uuid"

type StoryAlbumTagStoryAlbumsDTO struct {
	ID           uuid.UUID `json:"id"`
	TagId        uuid.UUID `json:"tag_id"`
	StoryAlbumId uuid.UUID `json:"story_album_id"`
}

