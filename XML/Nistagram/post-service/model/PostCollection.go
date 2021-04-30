package model

import "github.com/google/uuid"

type PostCollection struct {
	ID uuid.UUID `json:"id"`
	Title string `json:"title" gorm:"not null"`
	UserID uuid.UUID `json:"userID" gorm:"not null"`
	Posts []SinglePost `gorm:"many2many:post_collection_posts"`
}
