package model

import (
	"github.com/google/uuid"
)

type Content struct {
	ID uuid.UUID `json:"id"`
	Path string `json:"path" gorm:"not null"`
	Type ContentType `json:"type" gorm:"not null"`
}