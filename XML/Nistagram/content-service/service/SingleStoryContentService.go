package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
)

type SingleStoryContentService struct {
	Repo * repository.SingleStoryContentRepository
}

func (service * SingleStoryContentService) CreateSingleStoryContent(singleStoryContent *model.SingleStoryContent) error {
	err := service.Repo.CreateSingleStoryContent(singleStoryContent)
	if err != nil {
		return err
	}
	return nil
}
