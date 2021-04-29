package model

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID uuid.UUID `json: "id"`
	MessageContentID uuid.UUID `json:"messageContentID"`
	IsDisposable bool `json:"isDisposable" gorm:"not null"`
	CreationDate time.Time `json:"creationDate" gorm:"not null"`
	SenderUserID string `json:"senderUserID" gorm:"not null"`
	ReceiverUserID string `json:"receiverUserID" gorm:"not null"`
	IsDeleted bool `json:"isDeleted" gorm:"not null"`
}


