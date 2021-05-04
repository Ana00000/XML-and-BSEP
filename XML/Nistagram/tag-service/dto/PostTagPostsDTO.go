package dto

import "github.com/google/uuid"

type PostTagPostsDTO struct {
	PostTagId uuid.UUID `json:"post_tag_id"`
	PostId uuid.UUID `json:"post_id"`
}
