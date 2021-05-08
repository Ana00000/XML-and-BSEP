package service

import (
	"../model"
	"../repository"
)

type CommentTagCommentsService struct {
	Repo * repository.CommentTagCommentsRepository
}

func (service * CommentTagCommentsService) CreateCommentTagComments(commentTagComments *model.CommentTagComments) error {
	err := service.Repo.CreateCommentTagComments(commentTagComments)
	if err != nil {
		return err
	}
	return nil
}
