package dto

import "github.com/google/uuid"

type CommentTagCommentsDTO struct {
	CommentTagId uuid.UUID `json:"comment_tag_id"`
	CommentId uuid.UUID `json:"comment_id"`
}