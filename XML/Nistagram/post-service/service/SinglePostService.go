package service

import (
	"../model"
	"../repository"
)

type SinglePostService struct {
	Repo * repository.SinglePostRepository
}

func (service * SinglePostService) CreateSinglePost(singlePost *model.SinglePost) error {
	err := service.Repo.CreateSinglePost(singlePost)
	if err != nil {
		return err
	}
	return nil
}
