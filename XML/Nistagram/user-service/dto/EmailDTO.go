package dto

type EmailDTO struct {
	Email string `json:"email" validate:"required,email"`
}
