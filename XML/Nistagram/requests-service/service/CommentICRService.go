package service

import (
	"../model"
	"../repository"
)

type CommentICRService struct {
	Repo * repository.CommentICRRepository
}

func (service * CommentICRService) CreateCommentICR(commentICR *model.CommentICR) error {
	err := service.Repo.CreateCommentICR(commentICR)
	if err != nil {
		return err
	}
	return nil
}
