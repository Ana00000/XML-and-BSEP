package dto

import (
	"../model"
	"github.com/google/uuid"
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
	UserCategory model.UserCategory `json:"user_category"`
	OfficialDocumentPath string `json:"official_document_path"`
	SettingsId uuid.UUID `json:"settings_id"`
}