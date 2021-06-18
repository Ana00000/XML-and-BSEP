package dto

import "github.com/google/uuid"

type VerificationRequestDTO struct {
	UserId  uuid.UUID `json:"user_id" validate:"required"`
	FirstName              string                       `json:"first_name" validate:"required,alpha,min=2,max=20"`
	LastName               string                       `json:"last_name" validate:"required,alpha,min=2,max=35"`
	OfficialDocumentPath   string                       `json:"official_document_path"`
	RegisteredUserCategory string `json:"registered_user_category" validate:"required"`
}
