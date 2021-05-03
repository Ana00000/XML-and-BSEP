package model

import "github.com/google/uuid"

type StoryMessageSubstance struct{
	MessageSubstance
	StoryId uuid.UUID `json:"story_id" gorm:"not null"`
}