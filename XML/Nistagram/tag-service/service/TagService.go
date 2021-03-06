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

func (service *TagService) FindTagByName(name string) *model.Tag {
	tag := service.Repo.FindTagByName(name)
	return tag
}

func (service *TagService) FindTagByType(tagType string) *model.Tag {
	tag := service.Repo.FindTagByType(tagType)
	return tag
}

func (service *TagService) FindAll() []model.Tag {
	tags := service.Repo.FindAll()
	if tags != nil {
		return tags
	}
	return nil
}

func (service *TagService) FindAllHashTags() []model.Tag {
	tags := service.Repo.FindAllHashTags()
	if tags != nil {
		return tags
	}
	return nil
}

func (service * TagService) FindTagForId(tagId uuid.UUID) model.Tag{
	tag := service.Repo.FindTagForId(tagId)
	return tag
}