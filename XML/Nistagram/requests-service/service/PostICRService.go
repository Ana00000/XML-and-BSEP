package service

import (
	"../model"
	"../repository"
)

type PostICRService struct {
	Repo * repository.PostICRRepository
}

func (service * PostICRService) CreatePostICR(postICR *model.PostICR) error {
	service.Repo.CreatePostICR(postICR)
	return nil
}
