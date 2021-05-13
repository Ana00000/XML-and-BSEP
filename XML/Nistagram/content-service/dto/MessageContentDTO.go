package dto

import (
	"github.com/google/uuid"
)

type MessageContentDTO struct {
	Path string `json:"path"`
	Type string `json:"type"`
	MessageSubstanceId uuid.UUID `json:"message_substance_id"`
}
