package service

import (
	"../model"
	"../repository"
)

type CommentICRService struct {
	Repo * repository.CommentICRRepository
}

func (service * CommentICRService) CreateCommentICR(commentICR *model.CommentICR) error {
	service.Repo.CreateCommentICR(commentICR)
	return nil
}
