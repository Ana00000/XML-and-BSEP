package model

import "github.com/google/uuid"

type UserTag struct {
	Tag
	UserId uuid.UUID `json:"user_id" gorm:"not null"`
}
