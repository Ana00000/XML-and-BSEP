package model

import "github.com/google/uuid"

type CommentContent struct {
	Content
	 uuid.UUID `json:"comment_id" gorm:"not null"`
}