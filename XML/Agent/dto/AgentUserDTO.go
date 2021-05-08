package dto

import (
	"../model"
)

type AgentDTO struct {
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
}
