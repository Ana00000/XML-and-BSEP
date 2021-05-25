package dto

import (
	"github.com/google/uuid"
)

type ProductDTO struct {
	PicturePath string `json:"picturePath" validate:"required,file"`
	Amount int `json:"amount" validate:"required,numeric,gt=0"`
	Price float32 `json:"price" validate:"required,numeric,gt=0"`
	AgentUserID uuid.UUID `json:"agent_user_id" validate:"uuid"`
}
