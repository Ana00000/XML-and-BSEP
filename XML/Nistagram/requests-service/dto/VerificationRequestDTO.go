package dto

import model "../model"

type VerificationRequestDTO struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	OfficialDocumentPath string `json:"official_document_path"`
	UserCategory model.UserCategory `json:"user_category"`
}