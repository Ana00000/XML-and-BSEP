package model

import (
	messageContentPath "../../content-service/model"
	"github.com/google/uuid"
)


type MessageSubstance struct{
	ID uuid.UUID `json:"id"`
	Contents []messageContentPath.MessageContent `json:"contents" gorm:"foreignKey:MessageSubstanceId"`
	Text string `json:"text" gorm:"not null"`
}
