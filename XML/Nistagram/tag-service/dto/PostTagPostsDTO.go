package dto

import "github.com/google/uuid"

type PostTagPostsDTO struct {
	TagId uuid.UUID `json:"tag_id"`
	PostId uuid.UUID `json:"post_id"`
}
