package dto

import (
	"github.com/google/uuid"
)

type ProductDTO struct {
	PicturePath string `json:"picturePath" validate:"required,file"`
	Amount int `json:"amount" validate:"required"`
	Price float32 `json:"price" validate:"required"`
	AgentUserID uuid.UUID `json:"agent_user_id" validate:"required"`
}
