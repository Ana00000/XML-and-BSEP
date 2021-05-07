package model

import (
	"github.com/google/uuid"
)


type MessageSubstance struct{
	ID uuid.UUID `json:"id"`
	Text string `json:"text" gorm:"not null"`
}
