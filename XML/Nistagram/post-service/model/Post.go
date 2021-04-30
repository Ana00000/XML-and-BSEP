package model

import (
	postICRPath "../../requests-service/model"
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description" gorm:"not null"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	UserID uuid.UUID `json:"userID" gorm:"not null"`
	Activities []Activity `json:"activities" gorm:"foreignKey:PostID"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
	LocationID uuid.UUID `json:"locationID" gorm:"not null"`
	IsDeleted bool `json:"isDeleted" gorm:"not null"`
	PostICRs []postICRPath.PostICR `json:"postICRs" gorm:"foreignKey:PostId"`
}
