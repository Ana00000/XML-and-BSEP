package service

import (
	"../model"
	"../repository"
)

type CommentTagService struct {
	Repo * repository.CommentTagRepository
}

func (service * CommentTagService) CreateCommentTag(commentTag *model.CommentTag) error {
	service.Repo.CreateCommentTag(commentTag)
	return nil
}
