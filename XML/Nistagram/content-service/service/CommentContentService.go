package service

import (
	"../model"
	"../repository"
)

type CommentContentService struct {
	Repo * repository.CommentContentRepository
}

func (service * CommentContentService) CreateCommentContent(commentContent *model.CommentContent) error {
	service.Repo.CreateCommentContent(commentContent)
	return nil
}
