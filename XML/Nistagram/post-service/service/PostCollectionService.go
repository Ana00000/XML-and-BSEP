package service

import (
	"../model"
	"../repository"
)

type PostCollectionService struct {
	Repo * repository.PostCollectionRepository
}

func (service * PostCollectionService) CreatePostCollection(postCollection *model.PostCollection) error {
	err := service.Repo.CreatePostCollection(postCollection)
	if err != nil {
		return err
	}
	return nil
}
