package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/google/uuid"
)

type SingleStoryContentDTO struct {
	Path string `json:"path"`
	Type model.ContentType `json:"type"`
	SingleStoryId uuid.UUID `json:"single_story_id"`
}
