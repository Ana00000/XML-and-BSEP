package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/google/uuid"
)

type ProfileSettingsDTO struct{
	UserId uuid.UUID `json:"user_id"`
	UserVisibility model.UserVisibility `json:"user_visibility"`
	MessageApprovalType model.MessageApprovalType `json:"message_approval_type"`
}
