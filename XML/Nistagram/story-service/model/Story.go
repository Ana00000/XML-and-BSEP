package model

import (
	"github.com/google/uuid"
	"time"
)

type Story struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	UserId uuid.UUID `json:"userId" gorm:"not null"`
	LocationId uuid.UUID `json:"locationId"`
	IsDeleted bool `json:"isDeleted" gorm:"not null"`
	IsExpired bool `json:"isExpired" gorm:"not null"`
	Type StoryType `json:"type" gorm:"not null"`
}