package service


import (
	"../model"
	"../repository"
)

type TagService struct {
	Repo * repository.TagRepository
}

func (service * TagService) CreateTag(tag *model.Tag) error {
	service.Repo.CreateTag(tag)
	return nil
}
