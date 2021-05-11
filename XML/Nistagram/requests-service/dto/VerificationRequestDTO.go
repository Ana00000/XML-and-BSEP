package dto

import model "../model"

type VerificationRequestDTO struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	OfficialDocumentPath string `json:"official_document_path"`
	RegisteredUserCategory model.RegisteredUserCategory `json:"registered_user_category"`
}