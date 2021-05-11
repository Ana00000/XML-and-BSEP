package dto

type UserChangePasswordDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}
