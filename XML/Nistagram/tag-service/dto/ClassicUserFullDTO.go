package dto

import "github.com/google/uuid"

type ClassicUserFullDTO struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Gender string `json:"gender"`
	DateOfBirth string `json:"dateOfBirth"`
	Website string `json:"website"`
	Biography string `json:"biography"`
	Salt string `json:"salt"`
	IsConfirmed bool `json:"is_confirmed"`
	UserType string `json:"user_type"`
	IsDeleted bool `json:"is_deleted"`
}
