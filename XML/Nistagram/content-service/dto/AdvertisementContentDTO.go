package dto

import (
	"../model"
)

type AdvertisementContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	Link string `json:"link"`
}
