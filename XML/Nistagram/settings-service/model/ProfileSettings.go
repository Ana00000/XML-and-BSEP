package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettings struct {
	ID uuid.UUID `json: "id"`
	UserId uuid.UUID `json:"userId" gorm:"not null"`
	UserVisibility UserVisibility `json:"userVisibility" gorm:"not null"`
	MessageApprovalType MessageApprovalType `json:"messageApprovalType" gorm:"not null"`
	//MutedProfiles []userPath.RegisteredUser `json:"mutedProfiles" gorm:"many2many:profile_settings_muted_profiles"`
	//BlockedProfiles []userPath.RegisteredUser `json:"blockedProfiles" gorm:"many2many:profile_settings_blocked_profiles"`
	//ApprovedMessageProfiles []userPath.RegisteredUser `json:"approvedProfiles" gorm:"many2many:profile_settings_approved_profiles"`
	//RejectedMessageProfiles []userPath.RegisteredUser `json:"rejectedProfiles" gorm:"many2many:profile_settings_rejected_profiles"`

}

func(profileSettings * ProfileSettings) BeforeCreate(scope *gorm.DB) error {
	profileSettings.ID = uuid.New()
	return nil
}
