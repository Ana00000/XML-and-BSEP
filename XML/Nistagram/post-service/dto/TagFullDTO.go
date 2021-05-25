package dto

import (
	"github.com/google/uuid"
)

type TagFullDTO struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	TagType string `json:"tag_type"`
}
