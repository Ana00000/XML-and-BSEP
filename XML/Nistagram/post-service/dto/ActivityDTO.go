package dto

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
)

type ActivityDTO struct {
	PostID uuid.UUID `json:"postID"`
	UserID uuid.UUID `json:"userID"`
	LikedStatus model.LikedStatus `json:"likedStatus"`
	IsFavorite bool `json:"isFavorite"`
}

