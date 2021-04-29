package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Content struct {
	ID uuid.UUID `json: "id"`
	Path string `json:"path" gorm:"not null"`
	Type ContentType `json:"type" gorm:"not null"`
}

func(content * Content) BeforeCreate(scope *gorm.DB) error {
	content.ID = uuid.New()
	return nil
}