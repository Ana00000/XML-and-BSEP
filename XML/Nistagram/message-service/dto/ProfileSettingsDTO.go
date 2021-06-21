package dto
import (
	"github.com/google/uuid"
)

type ProfileSettingsDTO struct {
	UserId              uuid.UUID `json:"user_id" validate:"required"`
	UserVisibility      string    `json:"user_visibility" validate:"required"`
	MessageApprovalType string    `json:"message_approval_type" validate:"required"`
	IsPostTaggable      bool      `json:"is_post_taggable" validate:"required"`
	IsStoryTaggable     bool      `json:"is_story_taggable" validate:"required"`
	IsCommentTaggable   bool      `json:"is_comment_taggable" validate:"required"`
	LikesNotifications string `json:"likes_notifications" validate:"required"`
	CommentsNotifications string `json:"comments_notifications" validate:"required"`
	MessagesNotifications string `json:"messages_notifications" validate:"required"`
}
