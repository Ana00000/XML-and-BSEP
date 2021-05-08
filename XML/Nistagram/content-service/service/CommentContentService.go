package service

import (
	"../model"
	"../repository"
)

type CommentContentService struct {
	Repo * repository.CommentContentRepository
}

func (service * CommentContentService) CreateCommentContent(commentContent *model.CommentContent) error {
	err := service.Repo.CreateCommentContent(commentContent)
	if err != nil {
		return err
	}
	return nil
}
