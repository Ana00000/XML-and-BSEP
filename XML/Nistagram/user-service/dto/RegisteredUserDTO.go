package dto

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
)

type ClassicUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Gender model.Gender `json:"gender"`
	DateOfBirth string `json:"dateOfBirth"`
	Website string `json:"website"`
	Biography string `json:"biography"`
	ClassicUserCategory model.ClassicUserCategory `json:"user_category"`
	OfficialDocumentPath string `json:"official_document_path"`
	SettingsId uuid.UUID `json:"settings_id"`
}
