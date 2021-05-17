package dto

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
)

type ProfileSettingsDTO struct{
	UserId uuid.UUID `json:"user_id"`
	UserVisibility model.UserVisibility `json:"user_visibility"`
	MessageApprovalType model.MessageApprovalType `json:"message_approval_type"`
	IsPostTaggable bool `json:"is_post_taggable"`
	IsStoryTaggable bool `json:"is_story_taggable"`
	IsCommentTaggable bool `json:"is_comment_taggable"`
}
