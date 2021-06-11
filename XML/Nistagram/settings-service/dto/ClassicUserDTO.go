package dto

import "github.com/google/uuid"

type ClassicUserDTO struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Username    string    `json:"username" validate:"required,min=2,max=30"`
	Password    string    `json:"password" validate:"required,min=10,max=30"`
	Email       string    `json:"email" validate:"required,email"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
	FirstName   string    `json:"firstName" validate:"required,alpha,min=2,max=20"`
	LastName    string    `json:"lastName" validate:"required,alpha,min=2,max=35"`
	Gender      string    `json:"gender" validate:"oneof=MALE FEMALE OTHER"`
	DateOfBirth string    `json:"dateOfBirth" validate:"required"`
	Website     string    `json:"website" validate:"required"`
	Biography   string    `json:"biography" validate:"required"`
	Salt        string    `json:"salt" validate:"required"`
	IsConfirmed bool      `json:"is_confirmed" validate:"required"`
	UserType    string    `json:"user_type" validate:"required"`
	IsDeleted   bool      `json:"is_deleted" validate:"required"`
}
