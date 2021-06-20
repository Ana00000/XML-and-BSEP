package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationRequest struct {
	ID                     uuid.UUID              `json:"id"`
	FirstName              string                 `json:"first_name" gorm:"not null"`
	LastName               string                 `json:"last_name" gorm:"not null"`
	OfficialDocumentPath   string                 `json:"official_document_path" gorm:"not null"`
	RegisteredUserCategory RegisteredUserCategory `json:"registered_user_category" gorm:"not null"`
	VerificationRequestStatus RequestStatus `json:"verification_request_status" gorm:"not null"`
	UserId  uuid.UUID    `json:"user_id"`
}

func (verificationRequest *VerificationRequest) BeforeCreate(scope *gorm.DB) error {
	verificationRequest.ID = uuid.New()
	return nil
}
