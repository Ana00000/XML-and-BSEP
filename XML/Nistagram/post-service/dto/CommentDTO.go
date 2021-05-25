package dto

import (
	"github.com/google/uuid"
)

type CommentDTO struct {
	CreationDate string `json:"creation_date" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
	PostID uuid.UUID `json:"post_id" validate:"required"`
	Text string `json:"text" validate:"required,min=1,max=200"`
}