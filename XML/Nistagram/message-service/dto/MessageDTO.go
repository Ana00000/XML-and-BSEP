package dto

import (
	"github.com/google/uuid"
)

type MessageDTO struct {
	MessageContentID  uuid.UUID `json:"messageContentID"`
	IsDisposable bool `json:"isDisposable"`
	CreationDate string `json:"creationDate"`
	SenderUserID uuid.UUID `json:"senderUserID"`
	ReceiverUserID uuid.UUID `json:"receiverUserID"`
	IsDeleted bool `json:"isDeleted"`
}

