package dto

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
)

type LogInResponseDTO struct {
	ID uuid.UUID `json:"id"`
	Token string `json:"token"`
	UserType model.UserType `json:"userType"`
}
