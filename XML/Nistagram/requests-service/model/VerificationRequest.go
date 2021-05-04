package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationRequest struct {
	ID uuid.UUID `json:"id"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName string `json:"last_name" gorm:"not null"`
	OfficialDocumentPath string `json:"official_document_path" gorm:"not null"`
	UserCategoryValue UserCategory `json:"user_category_value" gorm:"not null"`
}

func(verificationRequest * VerificationRequest) BeforeCreate(scope *gorm.DB) error {
	verificationRequest.ID = uuid.New()
	return nil
}
