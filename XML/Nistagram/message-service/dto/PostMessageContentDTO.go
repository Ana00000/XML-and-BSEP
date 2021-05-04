package dto

import "github.com/google/uuid"

type PostMessageSubstanceDTO struct {
	Text string `json:"text"`
	PostId uuid.UUID `json:"post_id"`
}
