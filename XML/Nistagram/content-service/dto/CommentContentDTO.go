package dto

import (
	"github.com/google/uuid"
)

type CommentContentDTO struct {
	Path string `json:"path"`
	Type string `json:"type"`
	CommentId uuid.UUID `json:"comment_id"`
}
