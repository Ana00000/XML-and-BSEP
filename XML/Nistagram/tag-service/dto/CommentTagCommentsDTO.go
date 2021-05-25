package dto

import "github.com/google/uuid"

type CommentTagCommentsDTO struct {
	TagId uuid.UUID `json:"tag_id"`
	CommentId uuid.UUID `json:"comment_id"`
}