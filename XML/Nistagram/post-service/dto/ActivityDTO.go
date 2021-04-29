package dto

import (
	"github.com/google/uuid"
)

type ActivityDTO struct {
	PostID uuid.UUID `json:"postID"`
	UserID uuid.UUID `json:"userID"`
	Liked bool `json:"liked"`
	IsFavorite bool `json:"isFavorite"`
}

