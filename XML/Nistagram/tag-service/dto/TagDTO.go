package dto

import "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"

type TagDTO struct {
	Name string `json:"name"`
	TagType model.TagType `json:"gender"`
}
