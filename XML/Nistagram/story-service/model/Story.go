package model

import (
	requestsPath "../../requests-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Story struct {
	ID uuid.UUID `json: "id"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	UserId string `json:"userId" gorm:"not null"`
	LocationId uuid.UUID `json:"locationId" gorm:"not null"`
	IsDeleted bool `json:"isDeleted" gorm:"not null"`
	Type StoryType `json:"type" gorm:"not null"`
	StoryICRs []requestsPath.StoryICR `json:"storyICRs" gorm:"foreignKey:StoryId"`
}

func(story * Story) BeforeCreate(scope *gorm.DB) error {
	story.ID = uuid.New()
	return nil
}