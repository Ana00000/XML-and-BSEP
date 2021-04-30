package service

import (
	"../model"
	"../repository"
)

type SinglePostService struct {
	Repo * repository.SinglePostRepository
}

func (service * SinglePostService) CreateSinglePost(singlePost *model.SinglePost) error {
	service.Repo.CreateSinglePost(singlePost)
	return nil
}
