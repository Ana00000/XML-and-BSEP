package dto

import (
	"github.com/google/uuid"
)

type ProfileSettingsDTO struct {
	UserId uuid.UUID `json:"user_id"`
	UserVisibility string `json:"user_visibility"`
	MessageApprovalType string `json:"message_approval_type"`
	IsPostTaggable bool `json:"is_post_taggable"`
	IsStoryTaggable bool `json:"is_story_taggable"`
	IsCommentTaggable bool `json:"is_comment_taggable"`
	LikesNotifications string `json:"likes_notifications"`
	CommentsNotifications string `json:"comments_notifications"`
	MessagesNotifications string `json:"messages_notifications"`
}