package model

import "github.com/google/uuid"

type MessageContent struct {
	Content
	MessageSubstanceId uuid.UUID `json:"message_substance_id" gorm:"not null"`
}
