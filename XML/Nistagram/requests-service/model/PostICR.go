package model

import "github.com/google/uuid"

type PostICR struct {
	InappropriateContentRequest
	PostId uuid.UUID `json:"postId" gorm:"not null"`
}
