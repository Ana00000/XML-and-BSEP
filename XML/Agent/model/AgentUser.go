package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AgentUser struct {
	ID uuid.UUID `json:"id"`
	Username     string `json:"username" gorm:"not null"`
	Password     string `json:"password" gorm:"not null"`
	Email        string `json:"email" gorm:"not null"`
	PhoneNumber  string `json:"phoneNumber" gorm:"not null"`
	FirstName    string `json:"firstName" gorm:"not null"`
	LastName     string `json:"lastName" gorm:"not null"`
	Gender       Gender `json:"gender" gorm:"not null"`
	DateOfBirth  time.Time `json:"dateOfBirth" gorm:"not null"`
	Website      string `json:"website" gorm:"not null"`
	Biography    string `json:"biography" gorm:"not null"`
}

func(agentUser * AgentUser) BeforeCreate(scope *gorm.DB) error {
	agentUser.ID = uuid.New()
	return nil
}
