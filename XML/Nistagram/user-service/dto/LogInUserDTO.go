package dto

type LogInUserDTO struct {
	Username string `json:"username" validate:"required,min=2,max=30"`
	Password string `json:"password" validate:"required,min=10,max=30"`
	Question string `json:"question"`
	Answer string `json:"answer"`
}
