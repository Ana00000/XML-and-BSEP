package model

import "github.com/google/uuid"

type SinglePostContent struct {
	Content
	SinglePostId uuid.UUID `json:"single_post_id" gorm:"not null"`
}