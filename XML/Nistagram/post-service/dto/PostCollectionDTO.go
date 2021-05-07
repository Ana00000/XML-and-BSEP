package dto

import "github.com/google/uuid"

type PostCollectionDTO struct{
	Title string `json:"title"`
	UserID uuid.UUID `json:"userID"`
}
