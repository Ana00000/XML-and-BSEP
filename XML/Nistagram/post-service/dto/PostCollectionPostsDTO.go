package dto

import "github.com/google/uuid"

type PostCollectionPostsDTO struct {
	PostCollectionId uuid.UUID `json:"post_collection_id"`
	SinglePostId uuid.UUID `json:"single_post_id"`
}
