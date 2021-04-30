package model

import (
	storyPath "../../story-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID uuid.UUID `json: "id"`
	Longitude string `json:"longitude" gorm:"not null"`
	Latitude string `json:"latitude" gorm:"not null"`
	Country string `json:"country" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
	StreetName string `json:"streetName" gorm:"not null"`
	StreetNumber string `json:"streetNumber" gorm:"not null"`
	Stories []storyPath.Story `json:"stories" gorm:"foreignKey:LocationId"`
}

func(location * Location) BeforeCreate(scope *gorm.DB) error {
	location.ID = uuid.New()
	return nil
}