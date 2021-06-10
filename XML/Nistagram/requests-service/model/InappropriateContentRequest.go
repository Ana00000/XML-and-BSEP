package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InappropriateContentRequest struct {
	ID     uuid.UUID `json:"id"`
	Note   string    `json:"note" gorm:"not null"`
	UserId uuid.UUID `json:"userId" gorm:"not null"`
}

func (inappropriateContentRequest *InappropriateContentRequest) BeforeCreate(scope *gorm.DB) error {
	inappropriateContentRequest.ID = uuid.New()
	return nil
}
