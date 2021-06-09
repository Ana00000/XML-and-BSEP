package model

import "github.com/google/uuid"

type CommentICR struct {
	InappropriateContentRequest
	CommentId uuid.UUID `json:"commentId" gorm:"not null"`
}
