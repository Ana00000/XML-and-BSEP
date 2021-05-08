package service

import (
	"../model"
	"../repository"
)

type PostICRService struct {
	Repo * repository.PostICRRepository
}

func (service * PostICRService) CreatePostICR(postICR *model.PostICR) error {
	err := service.Repo.CreatePostICR(postICR)
	if err != nil {
		return err
	}
	return nil
}
