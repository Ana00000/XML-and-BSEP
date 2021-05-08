package dto

import (
	"github.com/google/uuid"
)

type ProductDTO struct {
	PicturePath string `json:"picturePath"`
	Amount int `json:"amount"`
	Price float32 `json:"price"`
	AgentUserID uuid.UUID `json:"agent_user_id"`
}
