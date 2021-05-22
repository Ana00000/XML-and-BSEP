package dto

import (
	"github.com/google/uuid"
)

type CommentDTO struct {
	CreationDate string `json:"creation_date"`
	UserID uuid.UUID `json:"user_id"`
	PostID uuid.UUID `json:"post_id"`
	Text string `json:"text"`
}