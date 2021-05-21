package dto

import "github.com/google/uuid"

type UserTagDTO struct {
	Name string `json:"name" validate:"required"`
	TagType string `json:"tag_type" validate:"USER_TAG"`
	UserId uuid.UUID `json:"user_id" validate:"required"`
}

