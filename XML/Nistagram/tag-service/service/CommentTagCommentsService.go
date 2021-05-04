package service

import (
	"../model"
	"../repository"
)

type CommentTagCommentsService struct {
	Repo * repository.CommentTagCommentsRepository
}

func (service * CommentTagCommentsService) CreateCommentTagComments(commentTagComments *model.CommentTagComments) error {
	service.Repo.CreateCommentTagComments(commentTagComments)
	return nil
}
