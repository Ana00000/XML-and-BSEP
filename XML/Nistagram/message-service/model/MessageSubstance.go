package model

import (
	"github.com/google/uuid"
)


type MessageSubstance struct{
	ID uuid.UUID `json:"id"`
	//Contents []messageContentPath.MessageContent `json:"contents" gorm:"foreignKey:MessageSubstanceId"`
	//MessageId uuid.UUID `json:"message_id" gorm:"not null"`
	Text string `json:"text" gorm:"not null"`
}
