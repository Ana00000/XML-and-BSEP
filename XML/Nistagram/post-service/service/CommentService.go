package service

import (
	"../model"
	"../repository"
)

type CommentService struct {
	Repo * repository.CommentRepository
}

func (service * CommentService) CreateComment(comment *model.Comment) error {
	err := service.Repo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}