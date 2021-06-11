package dto

import "github.com/google/uuid"

type CommentICRDTO struct {
	Note      string    `json:"note" validate:"required"`
	UserId    uuid.UUID `json:"userId" validate:"required"`
	CommentId uuid.UUID `json:"commentId" validate:"required"`
}
