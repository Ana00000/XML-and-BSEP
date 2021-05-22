package dto

import "github.com/google/uuid"

type PostTagPostsDTO struct {
	ID        uuid.UUID `json:"id"`
	PostTagId uuid.UUID `json:"post_tag_id"`
	PostId    uuid.UUID `json:"post_id"`
}
