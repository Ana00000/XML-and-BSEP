package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
)

type ContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
}