package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/google/uuid"
)

type CommentContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	CommentId uuid.UUID `json:"comment_id"`
}
