package service

import (
	"../model"
	"../repository"
)

type ContentService struct {
	Repo * repository.ContentRepository
}

func (service * ContentService) CreateContent(content *model.Content) error {
	service.Repo.CreateContent(content)
	return nil
}