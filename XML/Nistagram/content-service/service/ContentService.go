package service

import (
	"../model"
	"../repository"
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