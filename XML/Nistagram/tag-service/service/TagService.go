package service

import (
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
