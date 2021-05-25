package model

import "github.com/google/uuid"

type PostMessageSubstance struct{
	MessageSubstance
	PostId uuid.UUID `json:"post_id" gorm:"not null"`
}