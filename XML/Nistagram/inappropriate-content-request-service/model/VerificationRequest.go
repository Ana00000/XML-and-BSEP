package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationRequest struct {
	ID uuid.UUID `json: "id"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName string `json: "lastName" gorm:"not null"`
	OfficialDocumentPath string `json: "officialDocumentPath" gorm:"not null"`
}

func(verificationRequest * VerificationRequest) BeforeCreate(scope *gorm.DB) error {
	verificationRequest.ID = uuid.New()
	return nil
}
