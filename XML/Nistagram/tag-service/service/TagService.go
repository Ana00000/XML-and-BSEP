package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
)

type TagService struct {
	Repo * repository.TagRepository
}

func (service * TagService) CreateTag(tag *model.Tag) error {
	err := service.Repo.CreateTag(tag)
	if err != nil {
		return err
	}
	return nil
}


func (service *TagService) FindTagNameById(ID uuid.UUID) string {
	tag := service.Repo.FindTagNameById(ID)
	return tag
}


//FindAllPostsByTagId
func (service *TagService) FindTagIdByTagName(tagName string) *model.Tag {
	tag := service.Repo.FindTagIdByTagName(tagName)

	if tag == nil{
		return nil
	}
	return tag
}
