package dto

import (
	"../model"
	"time"
)

type RegisteredUserDTO struct {
	Username string `json:"username"`
	Password string `json: "password"`
	Email string `json: "email`
	PhoneNumber string `json:"phoneNumber""`
	FirstName string `json: "firstName"`
	LastName string `json: "lastName"`
	Gender model.Gender `json:"gender"`
	DateOfBirth time.Time `json: "dateOfBirth"`
	Website string `json: "website"`
	Biography string `json: "biography"`
}
