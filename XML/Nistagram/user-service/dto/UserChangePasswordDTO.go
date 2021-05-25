package dto

type UserChangePasswordDTO struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10,max=30"`
	ConfirmedPassword string `json:"confirmed_password" validate:"required,min=10,max=30"`
}
