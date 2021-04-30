package dto

import (
	"github.com/google/uuid"
)

type CommentDTO struct {
	ContentID uuid.UUID `json:"contentID"`
	CreationDate string `json:"creationDate"`
	UserID uuid.UUID `json:"userID"`
	PostID uuid.UUID `json:"postID"`
}