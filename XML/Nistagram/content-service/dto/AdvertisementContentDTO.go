package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/google/uuid"
)

type AdvertisementContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	Link string `json:"link"`
	AdvertisementId uuid.UUID `json:"advertisement_id"`
}
