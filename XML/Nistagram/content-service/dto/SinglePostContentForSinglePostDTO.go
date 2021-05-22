package dto

import "github.com/google/uuid"

type SinglePostContentForSinglePostDTO struct {
	ID uuid.UUID `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
	SinglePostId uuid.UUID `json:"single_post_id"`
}
