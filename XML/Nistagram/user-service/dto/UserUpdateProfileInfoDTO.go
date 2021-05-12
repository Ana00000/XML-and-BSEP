package dto

import (
	"github.com/google/uuid"
)

type UserUpdateProfileInfoDTO struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Gender string `json:"gender"`
	DateOfBirth string `json:"dateOfBirth"`
	Website string `json:"website"`
	Biography string `json:"biography"`
	UserType string `json:"userType"`
}
