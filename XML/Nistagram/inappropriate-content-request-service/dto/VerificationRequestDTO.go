package dto

type VerificationRequestDTO struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	OfficialDocumentPath string `json:"officialDocumentPath"`
}