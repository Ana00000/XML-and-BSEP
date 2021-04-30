package service

import (
	"../model"
	"../repository"
)

type PostCollectionService struct {
	Repo * repository.PostCollectionRepository
}

func (service * PostCollectionService) CreatePostCollection(postCollection *model.PostCollection) error {
	service.Repo.CreatePostCollection(postCollection)
	return nil
}
