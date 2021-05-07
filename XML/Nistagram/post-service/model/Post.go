package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description" gorm:"not null"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	UserID uuid.UUID `json:"userID" gorm:"not null"`
	LocationId uuid.UUID `json:"locationID" gorm:"not null"`
	IsDeleted bool `json:"isDeleted" gorm:"not null"`
}
