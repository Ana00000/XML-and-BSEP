package dto

import (
	"../model"
)

type ContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
}