package dto

import "github.com/google/uuid"

type CommentICRDTO struct {
	Note string `json:"note"`
	UserId uuid.UUID `json:"userId"`
	CommentId uuid.UUID `json:"commentId"`
}