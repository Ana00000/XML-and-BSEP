package dto

import "github.com/google/uuid"

type StoryAlbumTagStoryAlbumsDTO struct {
	StoryAlbumTagId uuid.UUID `json:"storyAlbumTagId"`
	StoryAlbumId uuid.UUID `json:"storyAlbumId"`
}
