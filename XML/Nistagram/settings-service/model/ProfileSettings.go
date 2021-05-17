package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettings struct {
	ID uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id" gorm:"not null"`
	UserVisibility UserVisibility `json:"user_visibility" gorm:"not null"`
	MessageApprovalType MessageApprovalType `json:"message_approval_type" gorm:"not null"`
	IsPostTaggable bool `json:"is_post_taggable" gorm:"not null"`
	IsStoryTaggable bool `json:"is_story_taggable" gorm:"not null"`
	IsCommentTaggable bool `json:"is_comment_taggable" gorm:"not null"`

}

func(profileSettings * ProfileSettings) BeforeCreate(scope *gorm.DB) error {
	profileSettings.ID = uuid.New()
	return nil
}
