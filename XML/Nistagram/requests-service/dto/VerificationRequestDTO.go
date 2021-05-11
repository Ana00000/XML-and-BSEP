package dto

import model "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"

type VerificationRequestDTO struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	OfficialDocumentPath string `json:"official_document_path"`
	RegisteredUserCategory model.RegisteredUserCategory `json:"registered_user_category"`
}