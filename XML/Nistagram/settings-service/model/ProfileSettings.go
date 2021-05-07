package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileSettings struct {
	ID uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId" gorm:"not null"`
	UserVisibility UserVisibility `json:"userVisibility" gorm:"not null"`
	MessageApprovalType MessageApprovalType `json:"messageApprovalType" gorm:"not null"`
}

func(profileSettings * ProfileSettings) BeforeCreate(scope *gorm.DB) error {
	profileSettings.ID = uuid.New()
	return nil
}
