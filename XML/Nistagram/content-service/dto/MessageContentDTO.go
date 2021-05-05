package dto

import (
	"../model"
	"github.com/google/uuid"
)

type MessageContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	MessageSubstanceId uuid.UUID `json:"message_substance_id"`
}
