package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID uuid.UUID `json:"id"`
	MessageSubstanceId uuid.UUID `json:"message_substance" gorm:"foreignKey:MessageId"`
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

