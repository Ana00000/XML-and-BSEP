package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/google/uuid"
)

type SinglePostContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	SinglePostId uuid.UUID `json:"single_post_id"`
}
