package service


import (
	"../model"
	"../repository"
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
