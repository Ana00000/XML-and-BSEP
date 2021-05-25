package model

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name" gorm:"not null"`
	TagType TagType `json:"tag_type" gorm:"not null"`
}
