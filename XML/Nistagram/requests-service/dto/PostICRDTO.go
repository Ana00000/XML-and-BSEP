package dto

import "github.com/google/uuid"

type PostICRDTO struct {
	Note   string    `json:"note" validate:"required"`
	UserId uuid.UUID `json:"userId" validate:"required"`
	PostId uuid.UUID `json:"postId" validate:"required"`
}
