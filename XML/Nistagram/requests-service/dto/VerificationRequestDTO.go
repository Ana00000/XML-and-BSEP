package dto

import model "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"

type VerificationRequestDTO struct {
	FirstName              string                       `json:"first_name" validate:"required,alpha,min=2,max=20"`
	LastName               string                       `json:"last_name" validate:"required,alpha,min=2,max=35"`
	OfficialDocumentPath   string                       `json:"official_document_path" validate:"required"`
	RegisteredUserCategory model.RegisteredUserCategory `json:"registered_user_category" validate:"required"`
}
