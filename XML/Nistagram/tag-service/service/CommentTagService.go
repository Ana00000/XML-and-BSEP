package service

import (
	"../model"
	"../repository"
)

type CommentTagService struct {
	Repo * repository.CommentTagRepository
}

func (service * CommentTagService) CreateCommentTag(commentTag *model.CommentTag) error {
	err := service.Repo.CreateCommentTag(commentTag)
	if err != nil {
		return err
	}
	return nil
}
