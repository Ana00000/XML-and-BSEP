package model

import "github.com/google/uuid"

type PostCollectionPosts struct {
	ID uuid.UUID `json:"id"`
	PostCollectionId uuid.UUID `json:"post_collection_id" gorm:"not null"`
	SinglePostId uuid.UUID `json:"single_post_id" gorm:"not null"`
}
