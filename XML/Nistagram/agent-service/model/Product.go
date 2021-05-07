package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID uuid.UUID `json:"id"`
	PicturePath string `json:"picturePath" gorm:"not null"`
	Amount int `json:"amount" gorm:"not null"`
	Price float32 `json:"price" gorm:"not null"`
}

func(product * Product) BeforeCreate(scope *gorm.DB) error {
	product.ID = uuid.New()
	return nil
}