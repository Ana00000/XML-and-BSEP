package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID uuid.UUID `json:"id"`
	MessageSubstanceID uuid.UUID `json:"messageContentID" gorm:"not null"`
	IsDisposable bool `json:"isDisposable" gorm:"not null"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	SenderUserID uuid.UUID `json:"senderUserID" gorm:"not null"`
	ReceiverUserID uuid.UUID `json:"receiverUserID" gorm:"not null"`
	IsDeleted bool `json:"isDeleted" gorm:"not null"`
}

func(message *Message) BeforeCreate(scope *gorm.DB) error {
	message.ID = uuid.New()
	return nil
}

