package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	ID uuid.UUID `json: "id"`
	Name string `json:"name" gorm:"not null"`
}

func(tag * Tag) BeforeCreate(scope *gorm.DB) error {
	tag.ID = uuid.New()
	return nil
}
