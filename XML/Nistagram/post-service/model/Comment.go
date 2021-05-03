package model

import (
	messageContentPath "../../content-service/model"
	commentICRPath "../../requests-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID uuid.UUID `json:"id"`
	Content messageContentPath.CommentContent `json:"content" gorm:"foreignKey:CommentId"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	UserID uuid.UUID `json:"userID" gorm:"not null"`
	PostID uuid.UUID `json:"postID" gorm:"not null"`
	CommentICRs []commentICRPath.CommentICR `json:"commentICRs" gorm:"foreignKey:CommentId"`
}

func (comment *Comment) BeforeCreate(scope *gorm.DB) error {
	if comment.ID == uuid.Nil {
		comment.ID = uuid.New()
	}
	return nil
}