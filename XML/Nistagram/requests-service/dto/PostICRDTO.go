package dto

import "github.com/google/uuid"

type PostICRDTO struct {
	Note string `json:"note"`
	UserId uuid.UUID `json:"userId"`
	PostId uuid.UUID `json:"postId"`
}