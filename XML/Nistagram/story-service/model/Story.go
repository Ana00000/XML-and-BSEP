package model

import (
	locationPath "../../location-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Story struct {
	ID uuid.UUID `json: "id"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	UserId string `json:"userId" gorm:"not null"`
	Location locationPath.Location `json:"location" gorm:"not null"`
	IsDeleted bool `json:"isDeleted" gorm:"not null"`
	Type StoryType `json:"type" gorm:"not null"`

}

func(story * Story) BeforeCreate(scope *gorm.DB) error {
	story.ID = uuid.New()
	return nil
}