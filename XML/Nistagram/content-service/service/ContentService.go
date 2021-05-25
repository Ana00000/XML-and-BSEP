package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
)

type ContentService struct {
	Repo * repository.ContentRepository
}

func (service * ContentService) CreateContent(content *model.Content) error {
	err := service.Repo.CreateContent(content)
	if err != nil {
		return err
	}
	return nil
}