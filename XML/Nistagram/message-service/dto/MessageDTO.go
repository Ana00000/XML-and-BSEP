package dto

import (
	"github.com/google/uuid"
)

type MessageDTO struct {
	MessageContentID  uuid.UUID `json:"messageContentID"`
	IsDisposable bool `json:"isDisposable"`
	CreationDate string `json:"creationDate"`
	SenderUserID string `json:"senderUserID"`
	ReceiverUserID string `json:"receiverUserID"`
	IsDeleted bool `json:"isDeleted"`
}

