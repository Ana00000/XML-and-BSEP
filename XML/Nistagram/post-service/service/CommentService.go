package service

import (
	"../model"
	"../repository"
)

type CommentService struct {
	Repo * repository.CommentRepository
}

func (service * CommentService) CreateComment(comment *model.Comment) error {
	service.Repo.CreateComment(comment)
	return nil
}