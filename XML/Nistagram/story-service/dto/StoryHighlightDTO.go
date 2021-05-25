package dto

import "github.com/google/uuid"

type StoryHighlightDTO struct {
	Title string `json:"title"`
	UserId uuid.UUID `json:"userId"`

}

