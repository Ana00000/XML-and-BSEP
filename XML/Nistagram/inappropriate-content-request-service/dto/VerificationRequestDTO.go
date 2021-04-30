package dto

import (
	"../../user-service/model"
)

type VerificationRequestDTO struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	OfficialDocumentPath string `json:"officialDocumentPath"`
	UserCategory model.UserCategory `json: "userCategory"`
}